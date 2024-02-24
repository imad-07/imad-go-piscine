package piscine

func IsNumeric(s string) bool {
	x := []rune(s)
	for i := 0; i < len(s); i++ {
		if !(x[i] >= '0' && x[i] <= '9') {
			return false
		}
	}
	return true
}
