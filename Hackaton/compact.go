package piscine

func Compact(ptr *[]string) int {
	count := 0
	array := []string{}
	for _, i := range *ptr {
		if i != "" {
			array = append(array, i)
			count++
		}
	}
	*ptr = array
	return count
}
