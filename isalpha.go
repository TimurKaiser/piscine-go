package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' || s[i] <= 0 || s[i] > 9 { // s[i] représente la positon, s[0] de Hello est H, de plus les minuscules sont un membres a part de l'ASCII
			return true
		}
	}
	return false
}
