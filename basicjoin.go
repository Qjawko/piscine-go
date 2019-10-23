package piscine

func BasicJoin(strs []string) string {
	var result string
	for _, str := range strs {
		result += str
	}
	return result
}
