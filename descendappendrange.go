package piscine

func DescendAppendRange(max, min int) []int {
	var rslt []int
	if max <= min {
		r := []int{}
		return r
	}
	for i := max; i > min; i-- {
		rslt = append(rslt, i)
	}
	return rslt
}
