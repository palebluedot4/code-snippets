package mathutil

import (
	"errors"
)

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
