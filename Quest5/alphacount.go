package piscine

func AlphaCount(s string) int {
	i := 0
	for _, char := range s { // _, permet de prendre uniquement les valeurs qui nous intéresses BriceGPT
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			i++
		}
	}
	return i
}
