package abs

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func AbsIEEE[T constraints.Signed | constraints.Float](x T) T {
	switch any(x).(type) {
	case float32, float64:
		if math.Signbit(float64(x)) {
			return -x
		}
		return x
	default:
		if x < 0 {
			return -x
		}
		return x
	}
}
