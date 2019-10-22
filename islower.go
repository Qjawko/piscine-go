package piscine

func IsLower(str string) bool {
	for _, ch := range []rune(str) {
		if !(IsRuneLowerCase(ch) || IsRuneUpperCase(ch)) {
			return false
		}
	}

	return true
}
