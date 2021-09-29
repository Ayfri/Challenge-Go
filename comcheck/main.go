package main

import (
	"fmt"
	"os"
)

func main() {
	for _, element := range os.Args[1:] {
		if element == "01" || element == "galaxy" || element == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
