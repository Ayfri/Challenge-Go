package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	SortWordArr(args)

	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func SortWordArr(a []string) {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
