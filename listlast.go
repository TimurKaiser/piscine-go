package piscine

func ListLast(l *List) interface{} {
	for l.Tail == nil {
		return 0
	}
	return l.Tail
}
