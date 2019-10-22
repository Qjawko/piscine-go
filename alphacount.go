package piscine

func AlphaCount(str string) int {
	count := 0
	for _, ch := range str {
		for i := 'A'; i <= 'Z'; i++ {
			if ch == i {
				count++
			}
		}

		for i := 'a'; i <= 'z'; i++ {
			if ch == i {
				count++
			}
		}
	}

	return count
}
