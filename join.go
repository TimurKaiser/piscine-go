package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, str := range strs {
		result = result + str
		if i < len(strs)-1 {
			result = result + sep
		}
	}
	return result
}

