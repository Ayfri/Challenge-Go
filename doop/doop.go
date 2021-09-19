package main

func doop(value int, operator string, other int) {
	switch operator {
	case "+":
		println(value + other)

	case "-":
		println(value - other)

	case "*":
		println(value * other)

	case "/":
		if value == 0 && other == value {
			println("No division by 0")
		} else {
			println(value / other)
		}

	case "%":
		if value == 0 && other == value {
			println("No modulo by 0")
		} else {
			println(value % other)
		}

	default:
		return
	}
}
