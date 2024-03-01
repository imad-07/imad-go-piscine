package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				a[i] = -1
				a[j] = -1
			}
		}
	}
	for x := 0; x < len(a); x++ {
		if a[x] != -1 {
			return a[x]
		}
	}
	return -1
}
