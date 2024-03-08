package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, e := range tab {
		if f(e) {
			x++
		}
	}
	return x
}
