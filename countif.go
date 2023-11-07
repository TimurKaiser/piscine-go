package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for _, i := range tab {
		if f(i) {
			return result
		}
	}
	return 0
}
