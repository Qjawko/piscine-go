package piscine

func Join(strs []string, sep string) string {
	var result string
	for i := 0; i < StrsLen(strs); i++ {
		if i+1 == StrsLen(strs) {
			result += strs[i]
		} else {
			result += strs[i] + sep
		}
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
