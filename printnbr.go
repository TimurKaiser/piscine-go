package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	printNbrRec(n)
}

func printNbrRec(n int) {
	if n/10 > 0 {
		printNbrRec(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
