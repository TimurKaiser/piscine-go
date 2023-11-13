package main

import (
	"fmt"
	"os"
	"strconv"
)

func isQuadA(x, y int) bool {
	if x <= 0 || y <= 0 {
		return false
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Utilisation: quadAchecker <largeur> <hauteur>")
		return
	}

	x, errX := strconv.Atoi(args[0])
	y, errY := strconv.Atoi(args[1])

	if errX != nil || errY != nil {
		fmt.Println("NON")
		return
	}

	if isQuadA(x, y) {
		fmt.Printf("[quadA] [%d] [%d]\n", x, y)
	} else {
		fmt.Println("NON")
	}
}
