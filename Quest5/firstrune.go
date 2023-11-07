package piscine

func FirstRune(s string) rune {
	runes := []rune(s) // commande vue dans quest3 exo strlen
	if len(runes) > 0 {
		return runes[0] // l'index commence par 0 et non par 1
	}
	return 0 // obligation de mettre un return quelque chose pour les accolades
}
