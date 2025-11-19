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

func AbsFloat[T constraints.Float](x T) T {
	return T(math.Abs(float64(x)))
}

func absFloatSignbit[T constraints.Float](x T) T {
	if math.Signbit(float64(x)) {
		return -x
	}
	return x
}

func AbsBitwiseSigned[T constraints.Signed](x T) T {
	const bitsPerByte = 8
	shift := unsafe.Sizeof(x)*bitsPerByte - 1
	mask := x >> shift
	return (x + mask) ^ mask
}

func AbsBitwiseFloat[T constraints.Float](x T) T {
	switch unsafe.Sizeof(x) {
	case 4:
		bits := *(*uint32)(unsafe.Pointer(&x)) &^ (1 << 31)
		return *(*T)(unsafe.Pointer(&bits))
	case 8:
		bits := *(*uint64)(unsafe.Pointer(&x)) &^ (1 << 63)
		return *(*T)(unsafe.Pointer(&bits))
	default:
		panic(fmt.Sprintf("abs: unsupported float size %d", unsafe.Sizeof(x)))
	}
}

func absFloatBitwiseBranch[T constraints.Float](x T) T {
	switch unsafe.Sizeof(x) {
	case 4:
		if *(*uint32)(unsafe.Pointer(&x))&(1<<31) != 0 {
			return -x
		}
		return x
	case 8:
		if *(*uint64)(unsafe.Pointer(&x))&(1<<63) != 0 {
			return -x
		}
		return x
	default:
		panic(fmt.Sprintf("abs: unsupported float size %d", unsafe.Sizeof(x)))
	}
}
