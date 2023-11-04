package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return 0
	}
	taille := max - min
	var result []int

	for i := 0; i < taille; i++ {
		result = append(result, min+i) // append sert à defiler et mettre les valeurs comme make
	}
	return result
}
