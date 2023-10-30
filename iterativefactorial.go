package piscine

func IterativeFactorial(nb int) int {
	i := 1
	if nb <= 0 || nb >= 25 {
		return 0
	}
	for k := 1; k < nb; k++ {
		i = i * k
		k++
	}
	return nb
}
