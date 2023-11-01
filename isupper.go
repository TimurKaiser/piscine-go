package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' { // s[i] représente la positon, s[0] de Hello est H, de plus les majuscules sont un membres a part de l'ASCII
			return false
		}
	}
	return true
}
