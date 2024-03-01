package piscine

func Rot14(s string) string {
	x := []rune(s)
	for i := 0; i < len(x); i++ {
		if (x[i] <= 'n' && x[i] >= 'a') || (x[i] <= 'N' && x[i] >= 'A') {
			x[i] += 14
		} else if (x[i] >= 'o' && x[i] <= 'z') || (x[i] >= 'O' && x[i] <= 'Z') {
			x[i] -= 13
			x[i] += 1
		}
	}
	return string(x)
}
