package piscine

func CollatzCountdown(start int) int {
	steps := 0

	if start <= 0 {
		return -1
	}
	for start != 1 { // ca continue tant que c'est pas egal a 1
		if start%2 == 0 { // en gros si le reste(%) de la division euclidienne est 0 c'est divisible par deux donc on divise par deux
			start = start / 2
		} else {
			start = start*3 + 1
		}
		steps++
	}
	return steps
}
