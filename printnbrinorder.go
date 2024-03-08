package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	x := [10]int{}
	for n != 0 {
		digit := n % 10
		x[digit]++
		n /= 10
	}
	for i := 0; i <= 9; i++ {
		for x[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			x[i]--
		}
	}
}
