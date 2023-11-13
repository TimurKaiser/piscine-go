package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil { // En gros si la liste est vide il faut faire une nouvelle chaine newNode de plus
		l.Head = newNode
		l.Tail = newNode
	} else { // si la liste n'est pas vide, donc ca ajoute le nouveau liste a la fin du noeud
		l.Tail.Next = newNode
		l.Tail = newNode
	}

}
