package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, len(a))

	for i, ans := range a {
		tab[i] = f(ans)
	}
	return tab
}
