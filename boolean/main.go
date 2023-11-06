package main

import (
	"os"
	"github.com/01-edu/z01"
)

func printStr(s string) {
	str := []rune(s)
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool{
	if nbr % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	param := os.Args
	if isEven(len(param)-1) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
