package main

import (
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune(10)
}
