package piscine

func IsAlpha(str string) bool {
	flag := false
	for _, r := range str {
		if IsRuneDigit(r) || IsRuneLowerCase(r) || IsRuneUpperCase(r) {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}
