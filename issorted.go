package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a) -1 ; i++ {
		if f(a[i-1], a[i]) > 0  && f(a[i],a[i+1]) < 0 {
			return false
		}
	}

	return true
}

