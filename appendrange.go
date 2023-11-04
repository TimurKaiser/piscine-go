package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil // nil c'est comme 0 mais ca pose pas de soucis c'est un cheat code pour pas mal d'exos ca print rien genre 
	}
	taille := max - min
	var result []int

	for i := 0; i < taille; i++ {
		result = append(result, min+i) // append sert à defiler et mettre les valeurs comme make
	}
	return result
}
