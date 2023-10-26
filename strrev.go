package piscine

func StrRevs(s string) string {
	runes := []rune(s)
	var runesinv []runes
	for i := len(runes) - 1; i >= 0; i-- {
		runesinv = append(runesinv, runes[i])
	}
	return string(runesinv)
}
