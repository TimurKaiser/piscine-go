package main

import (
	"fmt"
	"os"
	"strconv"
)

func isQuadB(x, y int) bool {
	if x <= 0 || y <= 0 {
		return false
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Utilisation: quadBchecker <largeur> <hauteur>")
		return
	}

	x, errX := strconv.Atoi(args[0])
	y, errY := strconv.Atoi(args[1])

	if errX != nil || errY != nil {
		fmt.Println("NON")
		return
	}

	if isQuadB(x, y) {
		fmt.Printf("[quadB] [%d] [%d]\n", x, y)
	} else {
		fmt.Println("NON")
	}
}
