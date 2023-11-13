package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <quad> <x> <y>")
		return
	}

	quadName := os.Args[1]
	x, errX := strconv.Atoi(os.Args[2])
	y, errY := strconv.Atoi(os.Args[3])

	if errX != nil || errY != nil {
		fmt.Println("Invalid input. Please provide valid integer values for x and y.")
		return
	}

	switch quadName {
	case "quadA":
		QuadA(x, y)
	case "quadB":
		QuadB(x, y)
	case "quadC":
		QuadC(x, y)
	case "quadD":
		QuadD(x, y)
	case "quadE":
		QuadE(x, y)
	default:
		fmt.Println("Invalid quad name. Available quads: quadA, quadB, quadC, quadD, quadE.")
	}
}

func PrintRune(r rune, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(r)
	}
}

func PrintLine(a, b, c rune, x int) {
	z01.PrintRune(a)
	if x > 1 {
		PrintRune(b, x-2)
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	PrintLine('o', '-', 'o', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('|', ' ', '|', x)
		}
		PrintLine('o', '-', 'o', x)
	}
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	PrintLine('/', '*', '\\', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('|', ' ', '|', x)
		}
		PrintLine('\\', '*', '/', x)
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	PrintLine('A', 'B', 'A', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('B', ' ', 'B', x)
		}
		PrintLine('C', 'B', 'C', x)
	}
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	PrintLine('A', 'B', 'C', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('B', ' ', 'B', x)
		}
		PrintLine('A', 'B', 'C', x)
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	PrintLine('A', 'B', 'C', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('B', ' ', 'B', x)
		}
		PrintLine('C', 'B', 'A', x)
	}
}
