package piscine // J'AI PRIS LE CODE DE QUELQU'UN DAUTRE MAIS LE MIEN NE FONTIONNE PAS TELLEMENT J'AI MIS DE // J'EN AI MARRE; MON DES EST STOCKER SUR DISCORD

func TrimAtoi(s string) int {
	neg := false
	empty := true
	result := 0
	for _, char := range s {
		if empty && !neg && char == '-' {
			neg = true
		} else if IsRuneDigit(char) {
			result *= 10
			result += int(char - 48)
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if neg {
			return -result
		} else {
			return result
		}
	}
}

func IsRuneDigit(i rune) bool {
	if i >= '0' && i <= '9' {
		return true
	}
	return false
}
