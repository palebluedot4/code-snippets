package mathutil

import (
	"fmt"
	"math/big"
)

func Factorial(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("factorial: argument must be non-negative, got %d", n)
	}
	return new(big.Int).MulRange(1, n), nil
}

func FactorialIterative(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("factorial: argument must be non-negative, got %d", n)
	}
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result, nil
}

func FactorialRecursive(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("factorial: argument must be non-negative, got %d", n)
	}
	return factorialRecursive(new(big.Int).SetInt64(n)), nil
}

func factorialRecursive(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(1)) <= 0 {
		return big.NewInt(1)
	}
	return new(big.Int).Mul(n, factorialRecursive(new(big.Int).Sub(n, big.NewInt(1))))
}

const (
	// 20! is the largest factorial that fits in uint64
	MaxUint64Factorial = 20
)

func FactorialUint64Iterative(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial: argument must be non-negative, got %d", n)
	}
	if n > MaxUint64Factorial {
		return 0, fmt.Errorf("factorial: %d! would overflow uint64", n)
	}
	result := uint64(1)
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	return result, nil
}

func FactorialUint64Recursive(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial: argument must be non-negative, got %d", n)
	}
	if n > MaxUint64Factorial {
		return 0, fmt.Errorf("factorial: %d! would overflow uint64", n)
	}
	return factorialUint64Recursive(uint64(n)), nil
}

func factorialUint64Recursive(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorialUint64Recursive(n-1)
}
