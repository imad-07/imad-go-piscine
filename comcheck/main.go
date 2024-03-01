package main

import (
	"fmt"
	"os"
)

func main() {
	arg := []string(os.Args[1:])
	var a, b, c string
	a = "01"
	b = "galaxy"
	c = "galaxy 01"
	for i := 0; i < len(arg); i++ {
		if arg[i] == a || arg[i] == b || arg[i] == c {
			fmt.Println("Alert!!!")
			break
		}
	}
}
