package piscine

func IsAlpha(str string) bool {
	for _, r := range str {
		if !(IsRuneDigit(r) || IsRuneLowerCase(r) || IsRuneUpperCase(r)) {
			return false
		}
	}
	return true
}
