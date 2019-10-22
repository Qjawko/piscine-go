package piscine

func IsPrintable(str string) bool {
	for range []rune(str) {
		if contains(str, '\n') ||
			contains(str, '\x00') ||
			contains(str, '\t') {
			return false
		}
	}

	return true
}

func contains(s string, r rune) bool {
	for _, ch := range []rune(s) {
		if ch == r {
			return true
		}
	}

	return false
}
