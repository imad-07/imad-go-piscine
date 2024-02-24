package piscine

func IsAlpha(s string) bool {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if !(x[i] >= 'a' && x[i] <= 'z' || x[i] >= 'A' && x[i] <= 'Z' || x[i] >= '0' && x[i] <= '9') {
			return false
		}
	}
	return true
}
