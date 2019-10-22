package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	runesToFind := []rune(toFind)
	n := RuneLen(runesToFind)

	switch {
	case n == 0:
		return 1
	case n == 1:
		return IndexRune(s, runes[n])
	}

	if RuneLen(runes) < RuneLen(runesToFind) {
		return -1
	}

	for i := range runes {
		if runesToFind[0] == runes[i] {
			flag := false
			for j := range runesToFind {
				if runes[i+j] == runesToFind[j] {
					flag = true
				} else {
					flag = false
				}
			}

			if flag {
				return i
			}
		}
	}

	return -1
}

func IndexRune(s string, r rune) int {
	for i, ch := range []rune(s) {
		if ch == r {
			return i
		}
	}

	return -1
}
