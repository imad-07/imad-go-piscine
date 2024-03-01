package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	result := 0
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] < tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
	result = tab[2]
	return result
}
