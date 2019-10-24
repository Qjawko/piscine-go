package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var result []int = make([]int, max-min)
	for i, j := min, 0; i < max; i++ {
		result[j] = i
		j++
	}

	return result
}
