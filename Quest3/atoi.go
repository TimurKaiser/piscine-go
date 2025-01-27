package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	for i, char := range s {
		if i == 0 && (char == '+' || char == '-') {
			if char == '-' {
				sign = -1
			}
			continue
		}
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return result * sign
}
