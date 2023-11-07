package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	var res []int
	for i := 0; i < len(a); i += 2 {
		res = append(res, f(a[i], a[i+1]))
	}
	for j := 1; j < len(res); j++ {
		if res[j-1] > res[j] {
			return false
		}
	}
	return true
}
