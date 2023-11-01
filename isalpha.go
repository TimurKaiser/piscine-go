package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
		if s[i] < 'A' || s[i] > 'Z' {
			return false
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
