package piscine

func MakeRange(min, max int) []int {
	var tab []int
	if min < max {
		tab = make([]int, max-min)
		for i := 0; i < len(tab); i++ {
			tab[i] = min + i
		}
		return tab
	} else {
		return tab
	}
}

