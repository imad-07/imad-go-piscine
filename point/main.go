package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printnbr(n int) {
	nb := '0'
	subnb := '0'
	for i := 0; i < n/10; i++ {
		nb++
	}
	z01.PrintRune(nb)
	for i := 1; i <= n%10; i++ {
		subnb++
	}
	z01.PrintRune(subnb)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	printStr("x = ")
	printnbr(points.x)
	printStr(", y = ")
	printnbr(points.y)
	z01.PrintRune('\n')
}
