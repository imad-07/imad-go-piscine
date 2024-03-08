package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func printstr01(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	printstr01("Door Opening...")
	ptrDoor.state = "OPEN"
}

func CloseDoor(ptrDoor *Door) {
	printstr01("Door Closing...")
	ptrDoor.state = "CLOSE"
}

func IsDoorOpen(ptrDoor *Door) bool {
	printstr01("is the Door opened ?")
	return ptrDoor.state == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	printstr01("is the Door closed ?")
	return ptrDoor.state == "CLOSE"
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
