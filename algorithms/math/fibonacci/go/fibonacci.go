package fibonacci

import "fmt"

const (
	// The 93rd is the largest fibonacci number that fits in uint64
	MaxUint64Fibonacci = 93
)

func FibonacciUint64Iterative(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("fibonacci: argument must be a non-negative integer, got %d", n)
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
		return 0, fmt.Errorf("fibonacci: argument must be a non-negative integer, got %d", n)
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
