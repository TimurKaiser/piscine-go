package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && (s[i] < '0' || s[i] > '9') { // il vaut mieux eviter les arbres de if, c'est mieux d'enchainer les ||
			return false // enfaite il faut utiliser && car avec || il suffit qu'une des conditions soit valide et le reste est ignoré c'est pas ce que l'on veut
		} // je sais pas pourquoi mon code marche pas sans les () chat gpt m'a dit d'utiliser des () pour des suites logiques abcd... 1234..
	}
	return true
}
