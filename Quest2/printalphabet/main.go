package main

import "github.com/01-edu/z01"

func main() {
	char := 'a'
	for i := 1; i < 27; i++ {
		z01.PrintRune(char)
		char++

		z01.PrintRune('\n')

	}
}
