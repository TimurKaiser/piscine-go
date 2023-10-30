package piscine

func RecursivePower(nb int, power int) int {
	if nb < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if nb < 0 {
		return -RecursivePower(-nb, power)
	}
	r := nb * RecursivePower(nb, power-1)
	return r
}