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
	for i :=2; i <= index; i++;
	j = j + i
}
