package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, e := range a {
		result[i] = f(e)
	}
	return (result)
}
