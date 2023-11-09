package piscine

func Abort(a, b, c, d, e int) int {
	// Créer un tableau avec les cinq entiers
	nums := [5]int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	median := nums[2]

	return median
}
