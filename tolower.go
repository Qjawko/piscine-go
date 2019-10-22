package piscine

func ToLower(s string) string {
	runes := []rune(s)

	for i := range runes {
		if IsRuneUpperCase(runes[i]) {
			runes[i] += 32
		}
	}

	return string(runes)
}

func IsRuneUpperCase(r rune) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if r == i {
			return true
		}
	}
	return false
}
