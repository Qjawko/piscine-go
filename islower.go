package piscine

func IsLower(str string) bool {
	for _, ch := range []rune(str) {
		if !IsRuneLowerCase(ch) {
			return false
		}
	}

	return true
}
