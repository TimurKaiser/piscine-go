package piscine

func ToUpper(s string) string {
	r := ""
	for j := 0; j < len(s); j++ {
		k := s[j]
		if k >= 'a' && k <= 'z' {
			k = k - 32 // pour passer d'une minuscule à une majuscule il faut -32 sur l'ASCII j'ai compté trois fois frr
			r = r + string(j)
		}
	}
	return r
}
