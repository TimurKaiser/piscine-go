package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if i < 0 || i > 127 {
			return false
		}
	}
	return true
}
