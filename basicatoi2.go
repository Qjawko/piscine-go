package piscine

func BasicAtoi2(s string) int {
	nm := 0

	for _, ch := range s {
		if !containsIn0to9(ch) {
			return 0
		}
		nm = nm*10 + charToInt(ch)
	}

	return nm
}

func containsIn0to9(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}

	return false
}
