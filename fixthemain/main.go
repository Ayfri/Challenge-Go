package main

import "github.com/01-edu/z01"

const (
	CLOSE = iota
	OPEN
)

type Door struct {
	state int
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return door.state == OPEN
}

func IsDoorClose(door *Door) bool {
	PrintStr("is the Door closed ?\n")
	return door.state == OPEN
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
	if door.state == OPEN {
		CloseDoor(door)
	}
}
