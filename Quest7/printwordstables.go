package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var str []rune
	for j := 0; j < len(a); j++ {
		str = []rune(a[j])
		for i := 0; i < len(str); i++ {
			z01.PrintRune(str[i])
		}
		z01.PrintRune('\n')
	}
}
