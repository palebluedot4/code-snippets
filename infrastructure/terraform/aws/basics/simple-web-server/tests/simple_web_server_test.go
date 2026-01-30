package tests

import (
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	httphelper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleWebServer(t *testing.T) {
	t.Parallel()
	const (
		wantPort = 8080
		wantBody = "Hello, World!"
	)
	opts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]any{
			"server_port": wantPort,
		},
	})
	defer terraform.Destroy(t, opts)
	_, err := terraform.InitAndApplyE(t, opts)
	require.NoError(t, err)

	t.Run("outputs", func(t *testing.T) {
		tests := []struct {
			name       string
			outputName string
			assertion  func(t *testing.T, val string)
		}{
			{
				name:       "instance id format",
				outputName: "instance_id",
				assertion: func(t *testing.T, val string) {
					pattern := regexp.MustCompile(`^i-[a-z0-9]+$`)
					assert.Regexp(t, pattern, val)
				},
			},
			{
				name:       "public url validity",
				outputName: "instance_public_url",
				assertion: func(t *testing.T, val string) {
					u, err := url.Parse(val)
					require.NoError(t, err)
					assert.Equal(t, "http", u.Scheme)
					assert.Equal(t, strconv.Itoa(wantPort), u.Port())
					host, _, err := net.SplitHostPort(u.Host)
					require.NoError(t, err)
					ip := net.ParseIP(host)
					if assert.NotNil(t, ip) {
						assert.True(t, ip.IsGlobalUnicast())
						assert.NotNil(t, ip.To4())
					}
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := terraform.Output(t, opts, tt.outputName)
				tt.assertion(t, got)
			})
		}
	})

	t.Run("http connectivity", func(t *testing.T) {
		publicURL := terraform.Output(t, opts, "instance_public_url")
		maxRetries := 12
		sleepBetweenRetries := 10 * time.Second
		httphelper.HttpGetWithRetryWithCustomValidation(
			t,
			publicURL,
			nil,
			maxRetries,
			sleepBetweenRetries,
			func(statusCode int, body string) bool {
				return statusCode == http.StatusOK && strings.Contains(body, wantBody)
			},
		)
	})
}
