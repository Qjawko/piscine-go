package piscine

func IsUpper(str string) bool {
	for _, ch := range []rune(str) {
		if !IsRuneUpperCase(ch) {
			return false
		}
	}

	return true
}
