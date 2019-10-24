package piscine

func AppendRange(min, max int) []int {
	if max >= min {
		return nil
	}

	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
