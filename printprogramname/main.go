package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	PrintStr(os.Args[0])
}

func PrintStr(s string) {
	lastSlashIndex := 0
	for index, char := range s {
		if char == '/' {
			lastSlashIndex = index
		}
	}
	for i := lastSlashIndex; i < len(s); i++ {
		char := rune(s[i])
		z01.PrintRune(char)
	}
}
