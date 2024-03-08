package piscine

func Split(str string) []string {
	menu := []string{}
	add := ""
	for _, i := range str {
		if i != ' ' {
			add = add + string(i)
		} else {
			menu = append(menu, add)
			add = ""
		}
	}
	if add != " " {
		menu = append(menu, add)
	}
	return menu
}

func ShoppingSummaryCounter(str string) map[string]int {
	items := Split(str)
	conter := make(map[string]int)
	for _, i := range items {
		conter[i]++
	}
	return conter
}
