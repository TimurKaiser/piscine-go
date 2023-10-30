package piscine

func IterativeFactorial(nb int) int {
	if nb < 20 && nb >= 0 {
		i := 1
		for j := nb; j > 0; i-- {
			i = i * j
		}
		return i
	} else {
		return 0
	}
}
