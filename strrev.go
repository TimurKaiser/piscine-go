package piscine

func StrRev(s string) string {
	runes := []rune(s)
	var runesinv []rune
	for i := len(runes) - 1; i >= 0; i-- {
		runesinv = append(runesinv, runes[i])
	}
	return string(runesinv)
}
