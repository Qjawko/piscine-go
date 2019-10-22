package piscine

func NRune(s string, n int) rune {
	if n > RuneLen([]rune(s)) || n < 0 {
		return '\x00'
	}

	return []rune(s)[n-1]
}

func RuneLen(runes []rune) int {
	count := 0
	for range runes {
		count++
	}

	return count
}
