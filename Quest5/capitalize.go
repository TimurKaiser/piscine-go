package piscine

func Capitalize(s string) string {
	arg := []rune(s)
	letter := true
	for i := 0; i < len(s); i++ {
		if Prim(arg[i]) == true && letter {
			if arg[i] >= 'a' && arg[i] <= 'z' {
				arg[i] = 'A' - 'a' + arg[i]
			}
			letter = false
		} else if arg[i] >= 'A' && arg[i] <= 'Z' {
			arg[i] = 'a' - 'A' + arg[i]
		} else if Prim(arg[i]) == false {
			letter = true
		}
	}
	return string(arg)
}

func Prim(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}
