package piscine

func ReverseMenuIndex(menu []string) []string {
	result := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		result[i] = menu[len(menu)-i-1]
	}
	return result
}
