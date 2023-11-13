package piscine

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return 0
	}
	return l.Tail.Data
}
