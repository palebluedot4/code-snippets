package mathutil

import (
	"fmt"
	"math/big"
)

func Fibonacci(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("fibonacci: argument must be non-negative, got %d", n)
	}
	if n <= 1 {
		return big.NewInt(n), nil
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b, nil
}

func FibonacciRecursive(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("fibonacci: argument must be non-negative, got %d", n)
	}
	return fibonacciRecursive(new(big.Int).SetInt64(n)), nil
}

func fibonacciRecursive(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(1)) <= 0 {
		return n
	}
	nMinus1 := new(big.Int).Sub(n, big.NewInt(1))
	nMinus2 := new(big.Int).Sub(n, big.NewInt(2))
	return new(big.Int).Add(fibonacciRecursive(nMinus1), fibonacciRecursive(nMinus2))
}

const (
	// The 93rd is the largest fibonacci number that fits in uint64.
	MaxUint64Fibonacci = 93
)

func FibonacciUint64Iterative(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("fibonacci: argument must be non-negative, got %d", n)
	}
	if n > MaxUint64Fibonacci {
		return 0, fmt.Errorf("fibonacci: %dth fibonacci number would overflow uint64", n)
	}
	if n <= 1 {
		return uint64(n), nil
	}
	a := uint64(0)
	b := uint64(1)
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

func FibonacciUint64Recursive(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("fibonacci: argument must be non-negative, got %d", n)
	}
	if n > MaxUint64Fibonacci {
		return 0, fmt.Errorf("fibonacci: %dth fibonacci number would overflow uint64", n)
	}
	return fibonacciUint64Recursive(uint64(n)), nil
}

func fibonacciUint64Recursive(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return fibonacciUint64Recursive(n-1) + fibonacciUint64Recursive(n-2)
}
