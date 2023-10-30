package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 1 {
		return 1
	}
	if index == 0 {
		return 0
	}
	i := 0
	j := 1
	for k :=2; k <= index; k++ {
		temp := i //utiliser temp en tant que variable temporaire pour stocker une valeur temporairement pour la donner a une autre variable askipGPT
		i = j
		j = temp + j //du coup ca prends la valeur initiale de i malgré les modif c'est trop bien
	}
	return j
}
