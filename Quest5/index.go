package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ { // ligne à revoir apres je commprends pas trop la
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1 // vaux mieux rajouter -1 pour tout les autres cas ça rebouche les accolades et reponds aux negatifs
}
