package main

import (
	"os"
	"piscine"
)

func doop(value int, operator string, other int) {
	switch operator {
	case "+":
		piscine.PrintNbr(value + other)

	case "-":
		piscine.PrintNbr(value - other)

	case "*":
		piscine.PrintNbr(value * other)

	case "/":
		if value == 0 && other == value {
			os.Stdout.WriteString("No division by 0")
		} else {
			piscine.PrintNbr(value / other)
		}

	case "%":
		if value == 0 && other == value {
			os.Stdout.WriteString("No modulo by 0")
		} else {
			piscine.PrintNbr(value % other)
		}

	default:
		return
	}
}
