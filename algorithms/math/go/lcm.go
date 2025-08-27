package mathutil

import (
	"errors"
	"math/big"
)

func LCM(a, b *big.Int) (*big.Int, error) {
	if a == nil || b == nil {
		return nil, errors.New("lcm: nil argument")
	}
	if a.Sign() == 0 || b.Sign() == 0 {
		return big.NewInt(0), nil
	}
	absA := new(big.Int).Abs(a)
	absB := new(big.Int).Abs(b)
	gcd := new(big.Int).GCD(nil, nil, absA, absB)
	res := new(big.Int).Div(absA, gcd)
	return res.Mul(res, absB), nil
}

func LCMInt64(a, b int64) (int64, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}
	absA := Abs(a)
	absB := Abs(b)
	gcd := GCDInt64(absA, absB)
	if absA > ((1<<63-1)/absB)*gcd {
		return 0, errors.New("lcm: result would overflow int64")
	}
	return (absA / gcd) * absB, nil
}
