package piscine

func LastRune(s string) rune {
	length := RuneLen([]rune(s))

	if length == 0 {
		return '\x00'
	}

	return []rune(s)[length-1]
}
