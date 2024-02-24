package piscine

func Prime(a int) bool {
	if a <= 0 || a == 1 {
		return false
	}
	if a == 2 || a == 3 {
		return true
	}
	if a%2 == 0 || a%3 == 0 {
		return false
	}
	i := 5
	for i*i <= a {
		if a%i == 0 || a%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func FindNextPrime(nb int) int {
	if Prime(nb) {
		return nb
	}
	for i := 1; ; i++ {
		if Prime(nb + i) {
			return (nb + i)
		}
	}
}
