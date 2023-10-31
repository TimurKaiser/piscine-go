package piscine

func Concat(str1 string, str2 string) string {
	i := 1
	j := 1
	for i := 1; i < len(str1); i++ {
		j = j + string(str1[i]) // i permet d'aller à l'index de la position, par exemple le str[0] de Hello est "H"
	}
	for i := 1; i < len(str2); i++ {
		j = j + string(str2[1])
	}
	return j
}
