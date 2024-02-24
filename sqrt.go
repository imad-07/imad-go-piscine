package piscine

func Sqrt(nb int) int {
	if nb == 1 || nb == 0 {
		return nb
	}
	for i := 0; i <= nb/2; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
