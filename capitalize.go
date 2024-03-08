package piscine

func Capitalize(s string) string {
	rslt := []rune(s)
	if len(rslt) > 0 && rslt[0] >= 'a' && rslt[0] <= 'z' {
		rslt[0] -= 32
	}
	for i := 1; i < len(rslt); i++ {
		if !(isAlphanumeric(rslt[i-1])) {
			if rslt[i] >= 'a' && rslt[i] <= 'z' {
				rslt[i] -= 32
			}
		} else if rslt[i] >= 'A' && rslt[i] <= 'Z' {
			rslt[i] += 32
		}
	}
	return string(rslt)
}

func isAlphanumeric(x rune) bool {
	if !(x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' || x >= '0' && x <= '9') {
		return false
	}
	return true
}
