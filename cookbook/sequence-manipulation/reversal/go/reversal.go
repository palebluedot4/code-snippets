package reversal

import "slices"

func ReverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func ReverseSlice[T any](s []T) []T {
	reversed := make([]T, len(s))
	for i := range s {
		reversed[i] = s[len(s)-1-i]
	}
	return reversed
}

func ReverseSliceInPlace[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
