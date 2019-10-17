package piscine

func BasicAtoi(s string) int {
	nm := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		nm = nm*10 + int(ch)
	}

	return nm
}
