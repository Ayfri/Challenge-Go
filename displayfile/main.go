package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(content)
}
