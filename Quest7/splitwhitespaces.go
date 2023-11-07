package piscine

func SplitWhiteSpaces(s string) []string {
	str := []rune(s)
	var word []rune
	var result []string
	j := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '\n' || str[i] == '\t' {
			if string(word) != "" {
				result = append(result, string(word))
				word = nil
				j = 0
			}
		} else {
			word = append(word, str[i])
			j++
			if i == len(str)-1 {
				result = append(result, string(word))
			}
		}
	}
	return result
}
