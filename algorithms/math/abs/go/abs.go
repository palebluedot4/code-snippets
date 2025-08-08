package abs

import "golang.org/x/exp/constraints"

func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
