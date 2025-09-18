package reversal

import "slices"

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

func reverseSliceManual[S ~[]E, E any](s S) S {
	reversed := make(S, len(s))
	for i := range s {
		reversed[i] = s[len(s)-1-i]
	}
	return reversed
}

