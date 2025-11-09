package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	sloggin "github.com/samber/slog-gin"
)

const (
	defaultAppPort = "8080"

	postgresPasswordPath = "/run/secrets/postgres_password"
	redisPasswordPath    = "/run/secrets/redis_password"

	httpReadTimeout       = 10 * time.Second
	httpReadHeaderTimeout = 3 * time.Second
	httpWriteTimeout      = 15 * time.Second
	httpIdleTimeout       = 60 * time.Second
	httpMaxHeaderBytes    = 1 << 20
	httpShutdownTimeout   = 10 * time.Second

	postgresConnectTimeout    = 5 * time.Second
	postgresHealthCheckPeriod = 30 * time.Second
	postgresMaxConns          = 20
	postgresMinConns          = 5
	postgresMaxConnLifetime   = 10 * time.Minute
	postgresMaxConnIdleTime   = 2 * time.Minute

	redisDialTimeout  = 5 * time.Second
	redisPoolTimeout  = 2 * time.Second
	redisPoolSize     = 100
	redisMinIdleConns = 10
	redisReadTimeout  = 1 * time.Second
	redisWriteTimeout = 1 * time.Second
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	postgresPassword, err := os.ReadFile(postgresPasswordPath)
	if err != nil {
		logger.Error("failed to read PostgreSQL password", "path", postgresPasswordPath, "error", err)
		os.Exit(1)
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		strings.TrimSpace(string(postgresPassword)),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	dbConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		logger.Error("failed to parse PostgreSQL config", "error", err)
		os.Exit(1)
	}
	dbConfig.ConnConfig.ConnectTimeout = postgresConnectTimeout
	dbConfig.HealthCheckPeriod = postgresHealthCheckPeriod
	dbConfig.MaxConns = postgresMaxConns
	dbConfig.MinConns = postgresMinConns
	dbConfig.MaxConnLifetime = postgresMaxConnLifetime
	dbConfig.MaxConnIdleTime = postgresMaxConnIdleTime
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		logger.Error("failed to connect to PostgreSQL", "error", err)
		os.Exit(1)
	}
	defer dbPool.Close()
	if err := dbPool.Ping(ctx); err != nil {
		logger.Error("failed to ping PostgreSQL", "error", err)
		os.Exit(1)
	}
	logger.Info("successfully connected to PostgreSQL")

	redisPassword, err := os.ReadFile(redisPasswordPath)
	if err != nil {
		logger.Warn("failed to read Redis password, continuing without Redis", "path", redisPasswordPath, "error", err)
	} else {
		redisClient := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s",
				os.Getenv("REDIS_HOST"),
				os.Getenv("REDIS_PORT"),
			),
			Password:     strings.TrimSpace(string(redisPassword)),
			DialTimeout:  redisDialTimeout,
			PoolTimeout:  redisPoolTimeout,
			PoolSize:     redisPoolSize,
			MinIdleConns: redisMinIdleConns,
			ReadTimeout:  redisReadTimeout,
			WriteTimeout: redisWriteTimeout,
		})
		defer func() {
			if redisClient != nil {
				if e := redisClient.Close(); e != nil {
					logger.Error("failed to close Redis client", "error", e)
				}
			}
		}()
		if e := redisClient.Ping(ctx).Err(); e != nil {
			logger.Warn("failed to ping Redis", "error", e)
			redisClient = nil
		} else {
			logger.Info("successfully connected to Redis")
		}
	}

	router := gin.New()
	router.Use(sloggin.NewWithConfig(logger, sloggin.Config{
		WithUserAgent: true,
		WithRequestID: true,
	}))
	router.Use(gin.Recovery())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().UTC().Format(time.RFC3339Nano),
		})
	})

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = defaultAppPort
	}
	server := &http.Server{
		Addr:              ":" + appPort,
		Handler:           router,
		ReadTimeout:       httpReadTimeout,
		ReadHeaderTimeout: httpReadHeaderTimeout,
		WriteTimeout:      httpWriteTimeout,
		IdleTimeout:       httpIdleTimeout,
		MaxHeaderBytes:    httpMaxHeaderBytes,
		ErrorLog:          slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("starting server", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		logger.Error("server startup failed", "error", err)
		os.Exit(1)
	case <-ctx.Done():
		logger.Info("shutting down server")
		stop()
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), httpShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
		os.Exit(1)
	}
	logger.Info("server shutdown complete")
}
