package piscine

import "strconv"

func BasicAtoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
