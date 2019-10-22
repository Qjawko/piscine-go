package piscine

func BasicAtoi(s string) int {
	nm := 0
	for _, ch := range s {
		nm = nm*10 + charToInt(ch)
	}

	return nm
}

func charToInt(ch rune) int {
	count := 0
	if ch < 48 && ch > 58 {
		return 0
	}

	for i := '1'; i <= ch; i++ {
		count++
	}

	return count
}
