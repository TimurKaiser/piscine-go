package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	if len(runes) > 0 {
		return runes[len(runes)-1] // vue que la len commence par 1 et l'index par 0, la longueur -1 = l'index
	}
	return 0 // obligation de return un truc pour les accolades
}
