package piscine

func IsLower(s string) bool {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if x[i] < 'a' || x[i] > 'z' {
			return false
		}
	}
	return true
}
