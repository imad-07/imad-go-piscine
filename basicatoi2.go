package piscine

func BasicAtoi2(s string) int {
	r := 0
	x := []rune(s)
	l := len(x)
	for i := 0; i < l; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			r = r*10 + int(s[i]-'0')
		} else if s[i] == ' ' {
			return 0
		} else {
			return 0
		}
	}
	return r
}
