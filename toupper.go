package piscine

func ToUpper(s string) string {
	runes := []rune(s)

	for i := range runes {
		if IsRuneLowerCase(runes[i]) {
			runes[i] -= 32
		}
	}

	return string(runes)
}

func IsRuneLowerCase(r rune) bool {
	for i := 'a'; i <= 'z'; i++ {
		if r == i {
			return true
		}
	}
	return false
}
