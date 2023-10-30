package piscine

func IterativeFactorial(nb int) int {
	r := 1
	i := 1
	if nb < 0 {
		return 0
	}
	for j := 0; j < nb; j++ {
		r = r * i
		i++
		if j > 20 {
			return 0
		}
	}
}
