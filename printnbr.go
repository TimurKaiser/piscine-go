package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune("0123456789")
	printed := false

	for divisor := 1000000000; divisor > 0; divisor /= 10 {
		digit := (n / divisor) % 10
		if digit != 0 || printed {
			z01.PrintRune(digits[digit])
			printed = true
		}
	}
}
