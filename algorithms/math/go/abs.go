package mathutil

import (
	"fmt"
	"math"
	"unsafe"

	"golang.org/x/exp/constraints"
)

const (
	bitsPerByte = 8
	sizeFloat32 = 4
	sizeFloat64 = 8
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
	signBit := unsafe.Sizeof(x)*bitsPerByte - 1
	mask := x >> signBit
	return (x + mask) ^ mask
}

func AbsBitwiseFloat[T constraints.Float](x T) T {
	switch unsafe.Sizeof(x) {
	case sizeFloat32:
		const signBit = sizeFloat32*bitsPerByte - 1
		bits := *(*uint32)(unsafe.Pointer(&x)) &^ (1 << signBit)
		return *(*T)(unsafe.Pointer(&bits))
	case sizeFloat64:
		const signBit = sizeFloat64*bitsPerByte - 1
		bits := *(*uint64)(unsafe.Pointer(&x)) &^ (1 << signBit)
		return *(*T)(unsafe.Pointer(&bits))
	default:
		panic(fmt.Sprintf("abs: unsupported float size %d", unsafe.Sizeof(x)))
	}
}

func absFloatBitwiseBranch[T constraints.Float](x T) T {
	switch unsafe.Sizeof(x) {
	case sizeFloat32:
		const signBit = sizeFloat32*bitsPerByte - 1
		if *(*uint32)(unsafe.Pointer(&x))&(1<<signBit) != 0 {
			return -x
		}
		return x
	case sizeFloat64:
		const signBit = sizeFloat64*bitsPerByte - 1
		if *(*uint64)(unsafe.Pointer(&x))&(1<<signBit) != 0 {
			return -x
		}
		return x
	default:
		panic(fmt.Sprintf("abs: unsupported float size %d", unsafe.Sizeof(x)))
	}
}
