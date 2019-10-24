package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var result []int = make([]int, min, max)
	return result
}
