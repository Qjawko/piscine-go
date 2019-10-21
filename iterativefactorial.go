package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	}

	result := 1
	for i := nb; i >= 1; i-- {
		result *= i
	}

	return result
}
