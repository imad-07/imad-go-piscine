package piscine

func IsPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}
	y := nb / 2
	for x := 2; x <= y; x++ {
		if nb%x == 0 {
			return false
		}
	}
	return true
}
