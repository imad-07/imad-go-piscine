package piscine

func Join(strs []string, sep string) string {
	var res string
	for _, i := range strs {
		if res == "" {
			res = i
		} else {
			res = res + sep + i
		}
	}
	return res
}
