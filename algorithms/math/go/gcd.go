package mathutil

func GCDInt64(a, b int64) int64 {
	a = Abs(a)
	b = Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
