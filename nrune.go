package piscine

func NRune(s string, n int) rune {
	runes := []rune(s) // quest3 exo strlen
	if n < 1 {
		return 0 // si la valeur est négatif
	}
	if n > len(s) {
		return 0 // si il demande la 8eme lettre alors qu'il y'en a 3, cela return 0
	}
	return runes[n-1] // n vas s'adapter à l'index
}
