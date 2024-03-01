package piscine

func ActiveBits(n int) int {
	if n == 0 {
		return 0
	}
	rslt := 0
	for n > 0 {
		x := n % 2
		rslt += x
		n /= 2
	}
	return rslt
}
