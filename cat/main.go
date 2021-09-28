package main

import (
	"fmt"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Print(string(content))
	}
}
