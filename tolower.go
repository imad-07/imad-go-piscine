package piscine

func ToLower(s string) string {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if x[i] >= 'A' && x[i] <= 'Z' {
			x[i] += 32
		}
	}
	return string(x)
}
