package piscine

func Concat(str1 string, str2 string) string {
	j := ""
	for i := 1; i < len(str1); i++ { // faire commencer les boucles par i := 0 pour alligner avec l'index
		j = j + string(str1[i]) // i permet d'aller à l'index de la position, par exemple le str[0] de Hello est "H"
	}
	for i := 1; i < len(str2); i++ {
		j = j + string(str2[1])
	}
	return j
}
// on fait j ;= "" pour remplir des characteres et non des chiffres comme i:=1