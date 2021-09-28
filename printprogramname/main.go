package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lastSlashIndex := 0
	programName := os.Args[0]
	for index, char := range programName {
		if char == '/' {
			lastSlashIndex = index
		}
	}

	for i := lastSlashIndex - 1; i < len(programName); i++ {
		char := rune(programName[i])
		z01.PrintRune(char)
	}
}
