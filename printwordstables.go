package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	results := make([][]rune, len(a))
	for i, str := range a {
		results[i] = []rune(str)
	}
	for i := 0; i < len(results); i++ {
		for j := 0; j < len(results[i]); j++ {
			z01.PrintRune(results[i][j])
		}
		z01.PrintRune('\n')
	}
}
