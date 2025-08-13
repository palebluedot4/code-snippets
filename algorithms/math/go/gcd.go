package mathutil

import (
	"errors"
	"math/big"
)

func GCD(a, b *big.Int) (*big.Int, error) {
	if a == nil || b == nil {
		return nil, errors.New("gcd: nil argument")
	}
	return new(big.Int).GCD(nil, nil, a, b), nil
}

func GCDInt64(a, b int64) int64 {
	absA := Abs(a)
	absB := Abs(b)
	for absB != 0 {
		absA, absB = absB, absA%absB
	}
	return absA
}
