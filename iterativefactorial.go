package piscine

func IterativeFactorial(nb int) int {
	i := 1
	j := 1
	if nb < 0 {
		return 0
	}
	for k := 0; k < nb; k++ {
		i = i * j
		j++
	}
	return nb
}
