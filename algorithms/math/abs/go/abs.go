package abs

import (
	"math"
	"unsafe"

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

func AbsBitwiseSigned[T constraints.Signed](x T) T {
	const bitsPerByte = 8
	signBit := unsafe.Sizeof(x)*bitsPerByte - 1
	mask := x >> signBit
	return (x + mask) ^ mask
}
