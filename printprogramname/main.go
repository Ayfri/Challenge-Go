package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lastSlashIndex := -1
	programName := os.Args[0]
	for index, char := range programName {
		if char == '/' {
			lastSlashIndex = index
		}
	}

	for i := lastSlashIndex; i < len(programName); i++ {
		char := rune(programName[i])
		z01.PrintRune(char)
	}
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
