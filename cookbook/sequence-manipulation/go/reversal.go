package sequencemanipulation

import (
	"iter"
	"slices"
)

func ReverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func ReverseSlice[S ~[]E, E any](s S) S {
	clone := slices.Clone(s)
	slices.Reverse(clone)
	return clone
}

func ReverseSliceInPlace[S ~[]E, E any](s S) {
	slices.Reverse(s)
}

func ReverseSliceWithIter[S ~[]E, E any](s S) S {
	if len(s) < 2 {
		return slices.Clone(s)
	}
	reverseIter := func(yield func(E) bool) {
		for _, v := range slices.Backward(s) {
			if !yield(v) {
				return
			}
		}
	}
	return slices.Collect(reverseIter)
}

func reverseSliceManual[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	reversed := make(S, len(s))
	for i := range s {
		reversed[i] = s[len(s)-1-i]
	}
	return reversed
}

func reverseSliceInPlaceManual[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseSliceManualWithIter[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	if len(s) < 2 {
		return append(S{}, s...)
	}
	reverseIter := newReverseIter(s)
	reversed := make(S, 0, len(s))
	for v := range reverseIter {
		reversed = append(reversed, v)
	}
	return reversed
}

func newReverseIter[S ~[]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}
