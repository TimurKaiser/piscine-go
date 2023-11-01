package piscine

func ToLower(s string) string {
	r := ""
	for j := 0; j < len(s); j++ {
		k := s[j]
		if k >= 'A' && k <= 'Z' {
			k = k + 32 // J EN AI MARRE DE COMPTER CE TABLEAU JSP OU EST MON ERREUR
		}
		r = r + string(k)
	}
	return r
}
