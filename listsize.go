package piscine

func ListSize(l *List) int {
	taille := 0
	noeud := l.Head

	for noeud != nil { // != ( différent de )
		taille++
		noeud = noeud.Next //.Next ( noeud suivant)
	}
	return taille
}
