package main

import (
	"os"
	
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	es := args[1:]

	for i := 0; i < len(es); i++ {
		for j := 0; j < len(es)-1; j++ {
			if es[j] > es[j+1] {
				es[j], es[j+1] = es[j+1], es[j]
			}
		}
	}

	for _, v := range es {
		for _, char := range v {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
