package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var result []int = make([]int, min, max)
	for i, j := min, 0; i < max; i++ {
		result[j] = i
	}

	return result
	return result
}
