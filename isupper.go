package piscine

func IsUpper(s string) bool {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if x[i] < 'A' || x[i] > 'Z' {
			return false
		}
	}
	return true
}
