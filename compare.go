package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	if a < b { // l'exo est simple
		return -1
	}
	return 0
}
