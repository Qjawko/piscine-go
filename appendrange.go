package piscine

func AppendRange(min, max int) []int {
	if min < 0 {
		return nil
	}

	if max >= min {
		return nil
	}

	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
