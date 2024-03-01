package piscine

func StringToIntSlice(str string) []int {
	var result []int
	slice := []rune(str)
	for _, i := range slice {
		result = append(result, int(i))
	}
	return result
}
