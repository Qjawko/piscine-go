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
	return r >= 'A' && r <= 'Z'
}
