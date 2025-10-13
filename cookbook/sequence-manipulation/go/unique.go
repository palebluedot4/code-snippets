package sequencemanipulation

import "slices"

func Unique[S ~[]E, E comparable](s S) S {
	if len(s) < 2 {
		return slices.Clone(s)
	}
	seen := make(map[E]struct{}, len(s))
	res := make(S, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}
