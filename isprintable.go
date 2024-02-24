package piscine

func IsPrintable(s string) bool {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if !(x[i] >= 32 && x[i] <= 127) {
			return false
		}
	}
	return true
}
