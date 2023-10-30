package piscine

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	result := 1
	temp := 1
	if nb < 0 {
		return 0
	}
	for i := 0; i < nb; i++ {
		result = result * temp
		temp++
	}
}
