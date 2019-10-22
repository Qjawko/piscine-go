package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return '\x00'
	}

	if n > RuneLen([]rune(s)) {
		return '\x00'
	}

	return []rune(s)[n-1]
}

func RuneLen(runes []rune) int {
	count := 1
	for range runes {
		count++
	}

	return count
}
