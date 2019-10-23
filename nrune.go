package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return '\x00'
	}

	if n > RuneLen([]rune(s)) {
		return '\x00'
	}

	runes := []rune(s)
	return runes[n-1]
}

func RuneLen(runes []rune) int {
	count := 0
	for range runes {
		count++
	}

	return count
}
