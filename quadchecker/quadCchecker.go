package main

import (
	"fmt"
	"os"
	"strconv"
)

func isQuadC(x, y int) bool {
	if x <= 0 || y <= 0 {
		return false
	}

	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Utilisation: quadCchecker <largeur> <hauteur>")
		return
	}

	x, errX := strconv.Atoi(args[0])
	y, errY := strconv.Atoi(args[1])

	if errX != nil || errY != nil {
		fmt.Println("NON")
		return
	}

	if isQuadC(x, y) {
		fmt.Printf("[quadC] [%d] [%d]\n", x, y)
	} else {
		fmt.Println("NON")
	}
}
