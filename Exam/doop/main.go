package main // si tu reuissis ca tu reuissis la piscine c'est quoi ce truc merci quentin

import "os"

func main() {
	param := os.Args[1:]
	var oper []rune
	var nb1 []rune
	var nb2 []rune
	nb1int := 0
	nb2int := 0
	result := 0
	neg := false
	inval := false
	var resultStr []rune
	if len(param) == 3 && IsNumeric(param[0]) && IsNumeric(param[2]) {
		oper = []rune(param[1])
		nb1 = []rune(param[0])
		nb2 = []rune(param[2])
		nb1int = RunToInt(nb1)
		nb2int = RunToInt(nb2)
		if !(IfOverflow(nb1int)) && !(IfOverflow(nb2int)) {
			if oper[0] == '+' {
				result = nb1int + nb2int
			} else if oper[0] == '-' {
				result = nb1int - nb2int
			} else if oper[0] == '*' {
				result = nb1int * nb2int
			} else if oper[0] == '/' {
				if nb2int != 0 {
					result = nb1int / nb2int
				} else {
					inval = true
					os.Stdout.Write([]byte("No division by 0\n"))
				}
			} else if oper[0] == '%' {
				if nb2int != 0 {
					result = nb1int % nb2int
				} else {
					inval = true
					os.Stdout.Write([]byte("No modulo by 0\n"))
				}
			} else {
				inval = true
			}
			if !(inval) {
				if result < 0 {
					result *= -1
					neg = true
				}
				if result == 0 {
					resultStr = append(resultStr, '0')
				} else {
					for i := 0; result > 0; i++ {
						resultStr = append([]rune{rune(result%10) + 48}, resultStr...)
						result = result / 10
					}
				}
				if neg && resultStr != nil {
					resultStr = append([]rune{'-'}, resultStr...)
				}
				if resultStr != nil {
					os.Stdout.Write([]byte(string(resultStr)))
					os.Stdout.Write([]byte("\n"))
				}
			}
		}
	}
}

func IsNumeric(s string) bool {
	i := 0
	runes := []rune(s)
	for _, r := range s {
		_ = r
		if !(runes[i] >= '0' && runes[i] <= '9') && !(runes[i] == '-') {
			return false
		}
		i++
	}
	return true
}

func RunToInt(a []rune) int {
	decimal := 1
	result := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '-' {
			result *= -1
		} else {
			result += int(a[i]-48) * decimal
			decimal *= 10
		}
	}
	return result
}

func IfOverflow(a int) bool {
	testPos := 2147483647
	testNeg := -2147483648
	if testPos > a && testNeg < a {
		return false
	}
	return true
}
