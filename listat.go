package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	noeud := l
	for i := 0; i < pos; i++ {
		if noeud == nil {
			return nil
		}
		noeud = noeud.Next
	}
	return noeud
}
