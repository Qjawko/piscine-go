package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	length := RuneLen(runes)

	if IsRuneLowerCase(runes[0]) {
		runes[0] -= 32
	}

	for i := 1; i < length-1; i++ {
		if IsRuneUpperCase(runes[i]) {
			if IsRuneUpperCase(runes[i-1]) ||
				IsRuneLowerCase(runes[i-1]) ||
				IsRuneDigit(runes[i-1]) {
				runes[i] += 32
			}
		} else if IsRuneLowerCase(runes[i]) {
			if IsRuneUpperCase(runes[i-1]) ||
				IsRuneLowerCase(runes[i-1]) ||
				IsRuneDigit(runes[i-1]) {
				continue
			} else {
				runes[i] -= 32
			}
		}
	}

	return string(runes)
}

func IsRuneDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
