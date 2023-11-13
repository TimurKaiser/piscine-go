

package piscine

func ListAt(l *NodeL, pos int) *NodeL {
    noeud := 0

    for i := 0; i < pos; i++ {
        if noeud == nil {
            return pos
        }
        noeud = noeud.Next
    }

    return noeud
}