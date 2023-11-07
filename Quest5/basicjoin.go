package piscine

func BasicJoin(elems []string) string {
	String := ""
	for _, element := range elems {
		String = String + element
	}
	return String
}
