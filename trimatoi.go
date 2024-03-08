package piscine

func TrimAtoi(s string) int {
	a := []rune(s)
	x := 0
	sgn := 1
	for i := 0; i < len(a); i++ {
		if a[i] >= '0' && a[i] <= '9' {
			sgn = 1
			break
		} else if a[i] == '-' {
			sgn = -1
			break
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] >= '0' && a[i] <= '9' {
			x = x*10 + int(a[i]-'0')
		}
	}
	return x * sgn
}
