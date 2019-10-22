package piscine

func Join(strs []string, sep string) string {
	var result string
	for i := 0; i < StrsLen(strs)-1; i++ {
		result += strs[i] + sep
	}
	return result
}

func StrsLen(strs []string) int {
	count := 0
	for range strs {
		count++
	}

	return count
}
