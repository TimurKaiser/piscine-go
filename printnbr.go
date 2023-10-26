package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	str := strconv.Itoa(n)

	for _, char := range str {
		z01.PrintRune(char)
	}
}
