package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])
	for i := 2; i < len(arg); i++ {
		z01.PrintRune(arg[i])
	}
	z01.PrintRune('\n')
}
