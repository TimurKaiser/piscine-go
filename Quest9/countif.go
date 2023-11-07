package piscine

func CountIf(f func(string) bool, tab []string) int {
	r := 0
	for _, i := range tab {
		if f(i) == true {
			r++
		}
	}
	return r
}
