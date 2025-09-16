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

func ReverseSliceInPlace[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
