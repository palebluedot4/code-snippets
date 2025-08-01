package factorial

import "fmt"

const (
	// 20! is the largest factorial that fits in uint64
	MaxUint64Factorial = 20
)

func FactorialUint64Recursive(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial: argument must be a non-negative integer, got %d", n)
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
