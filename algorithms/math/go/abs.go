package mathutil

import (
	"fmt"
	"math"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func AbsSigned[T constraints.Signed](x T) T {
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
	shift := unsafe.Sizeof(x)*bitsPerByte - 1
	mask := x >> shift
	return (x + mask) ^ mask
}

func AbsBitwiseFloat[T constraints.Float](x T) T {
	switch any(x).(type) {
	case float32:
		bits := *(*uint32)(unsafe.Pointer(&x)) &^ (1 << 31)
		return *(*T)(unsafe.Pointer(&bits))
	case float64:
		bits := *(*uint64)(unsafe.Pointer(&x)) &^ (1 << 63)
		return *(*T)(unsafe.Pointer(&bits))
	default:
		panic(fmt.Sprintf("abs: unexpected type %T", x))
	}
}

func AbsIEEEBits[T constraints.Signed | constraints.Float](x T) T {
	switch any(x).(type) {
	case float32:
		if *(*uint32)(unsafe.Pointer(&x))&(1<<31) != 0 {
			return -x
		}
		return x
	case float64:
		if *(*uint64)(unsafe.Pointer(&x))&(1<<63) != 0 {
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
