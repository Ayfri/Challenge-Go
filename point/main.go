package point

import (
	"fmt"
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
	result := "x = " + ConvertNbr(points.x) + "y = " + ConvertNbr(points.y) + "\n"
	for _, char := range result {
		z01.PrintRune(char)
	}
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

	if nbr == 0 {
		str = "0"
	}

	return str
}
