package piscine

func AppendRange(min, max int) []int {
	if min > max {
		return nil
	}
	var x []int
	for i := min; i < max; i++ {
		x = append(x, i)
	}
	return x
}
