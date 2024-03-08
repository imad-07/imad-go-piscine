package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	var y rune
	if n <= 0 {
		return 0
	} else if n > len(x) {
		return 0
	} else {
		y = x[n-1]
	}
	return y
}
