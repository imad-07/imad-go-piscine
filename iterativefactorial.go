package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 21 {
		return 0
	}
	rslt := 1
	for i := nb; i >= 1; i-- {
		rslt = rslt * i
	}
	return (rslt)
}
