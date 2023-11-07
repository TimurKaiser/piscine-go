package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nameProgram := os.Args[0]
	str := []rune(nameProgram)

	for i := 2; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
	z01.PrintRune(10)
}
