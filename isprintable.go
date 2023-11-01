package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if i < 32 || i > 126 { // dans l'ASCII les char sont de 32 à 126
			return false
		}
	}
	return true
}
