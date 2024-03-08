package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	nonZeroCount := 0
	j := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[j] = slice[i]
			j++
			nonZeroCount++
		}
	}

	*ptr = slice[:nonZeroCount]
	return nonZeroCount
}
