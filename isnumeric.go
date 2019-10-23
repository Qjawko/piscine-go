package piscine

func IsNumeric(str string) bool {
	for _, r := range str {
		if !IsRuneDigit(r) {
			return false
		}
	}
	return true
}
