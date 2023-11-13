package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintRune(r rune, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(r))
	}
}

func PrintLine(a, b, c rune, x int) {
	fmt.Print(string(a))
	if x > 1 {
		PrintRune(b, x-2)
		fmt.Print(string(c))
	}
	fmt.Println()
}

func QuadA(x, y int) {
	PrintLine('o', '-', 'o', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('|', ' ', '|', x)
		}
		PrintLine('o', '-', 'o', x)
	}
}

func QuadB(x, y int) {
	PrintLine('/', '*', '\\', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('|', ' ', '|', x)
		}
		PrintLine('\\', '*', '/', x)
	}
}

func QuadC(x, y int) {
	PrintLine('A', 'B', 'A', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('B', ' ', 'B', x)
		}
		PrintLine('C', 'B', 'C', x)
	}
}

func QuadD(x, y int) {
	PrintLine('A', 'B', 'C', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('B', ' ', 'B', x)
		}
		PrintLine('A', 'B', 'C', x)
	}
}

func QuadE(x, y int) {
	PrintLine('A', 'B', 'C', x)
	if y > 1 {
		for i := 2; i < y; i++ {
			PrintLine('B', ' ', 'B', x)
		}
		PrintLine('C', 'B', 'A', x)
	}
}

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

	// Call the specified quad
	result := getQuadOutput(quadName, x, y)
	fmt.Println(result)

	// Check if other quads produce the same visual result
	sameVisualResults := FindSameVisualResults(quadName, x, y)
	if len(sameVisualResults) > 0 {
		fmt.Printf("%s | go run .\n", strings.Join(sameVisualResults, " || "))
	} else {
		fmt.Println("No visually similar quads found.")
	}
}

func FindSameVisualResults(quadName string, x, y int) []string {
	var result []string

	for _, q := range []string{"quadA", "quadB", "quadC", "quadD", "quadE"} {
		if q == quadName {
			continue
		}

		r, w := getQuadOutput(q, x, y), getQuadOutput(quadName, x, y)
		if r == w {
			result = append(result, fmt.Sprintf("[%s] [%d] [%d]", q, x, y))
		}
	}

	return result
}

func getQuadOutput(quadName string, x, y int) string {
	var result strings.Builder
	w := &result

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
	}

	return result.String()
}
