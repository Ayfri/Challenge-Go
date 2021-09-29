package main

import "piscine"

func main() {
	strings := []string{"a", "b", "c", "d", "e", "", "f", "g", "", "", "", "", "", ""}
	piscine.Compact(&strings)
}
