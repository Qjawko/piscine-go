package piscine

func Join(strs []string, sep string) string {
	var result string
	for _, str := range strs {
		result += str + sep
	}
	return result
}
