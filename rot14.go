package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	arr := []rune(s)
	var result string

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			if arr[i] >= 'm' {
				arr[i] -= 12
			} else {
				arr[i] += 14
			}
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			if arr[i] >= 'M' {
				arr[i] -= 12
			} else {
				arr[i] += 14
			}
		}
		result += string(arr[i])
	}
	return result
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
