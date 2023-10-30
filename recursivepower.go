package piscine

func RecursivePower(nb int, power int) int {
	r := 1
	if nb < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for i := 0; i < power; i++ {
		r = nb * power
	}
	return r
}
