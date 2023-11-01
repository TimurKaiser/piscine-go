package piscine

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '1' || s[i] > '9'{
			return false
		}
	}
	return true
}
