package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n > 0 {
		var nb []int
		for i3 := 0; n > 0; i3++ {
			nb = append(nb, n%10)
			n = n / 10
		}
		if len(nb) == 1 {
			z01.PrintRune(rune(nb[0] + 48))
		} else {
			bol := true // code à revoir car pris à quelqu'un d'autre pour avanver cause de retard, j'étais bloqué sur l'exo, à retravailler
			temp := 0
			for bol == true {
				bol = false
				for i := 1; i < len(nb); i++ {
					if nb[i-1] > nb[i] {
						temp = nb[i-1]
						nb[i-1] = nb[i]
						nb[i] = temp
						bol = true
					}
				}
			}
			for i2 := 0; i2 < len(nb); i2++ {
				z01.PrintRune(rune(nb[i2] + 48))
			}
		}
	} else if n == 0 {
		z01.PrintRune(rune(n + 48))
	}
}
