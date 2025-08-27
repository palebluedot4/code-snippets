package mathutil

import "golang.org/x/exp/constraints"

func IsEven[T constraints.Integer](x T) bool {
	return x%2 == 0
}

func IsEvenBitwise[T constraints.Integer](x T) bool {
	return x&1 == 0
}
