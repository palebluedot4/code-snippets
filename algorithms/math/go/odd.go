package mathutil

import "golang.org/x/exp/constraints"

func IsOdd[T constraints.Integer](x T) bool {
	return x%2 != 0
}

func IsOddBitwise[T constraints.Integer](x T) bool {
	return x&1 != 0
}
