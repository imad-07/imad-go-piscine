package piscine

func BasicAtoi(s string) int {
	x := []rune(s)
	r := 0
	for i := 0; i < len(x); i++ {
		y := (x[i] - '0')
		r = r*10 + int(y)
	}
	return r
}
