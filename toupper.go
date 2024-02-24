package piscine

func ToUpper(s string) string {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if x[i] >= 'a' && x[i] <= 'z' {
			x[i] -= 32
		}
	}
	return string(x)
}
