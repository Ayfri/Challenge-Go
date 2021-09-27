package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	PrintStr(os.Args[0])
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
