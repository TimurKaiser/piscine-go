package piscine

func RecursivePower(nb int, power int) int {
	r := 1
	if nb < 0 {
		return 0
	}
	return r == nb*power*power
}
