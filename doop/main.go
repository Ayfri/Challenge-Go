package main

import (
	"os"
)

func doop(value int, operator string, other int) {
	switch operator {
	case "+":
		printInt(value + other)

	case "-":
		printInt(value - other)

	case "*":
		printInt(value * other)

	case "/":
		if value == 0 && other == value {
			os.Stdout.WriteString("No division by 0")
		} else {
			printInt(value / other)
		}

	case "%":
		if value == 0 && other == value {
			os.Stdout.WriteString("No modulo by 0")
		} else {
			printInt(value % other)
		}

	default:
		return
	}
}

func printInt(nbr int) {
	os.Stdout.WriteString(ConvertNbr(nbr))
}

func ConvertNbr(nbr int) string {
	str := ""
	isNegative := nbr < 0

	for i, remainder := nbr, 0; i != 0; i /= 10 {
		remainder = i % 10
		if isNegative {
			remainder = -remainder
		}
		char := rune(remainder) + rune('0')
		str += string(char)
	}

	if isNegative {
		str += "-"
	} else if nbr == 0 {
		str = "0"
	}

	return Reverse(str)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
