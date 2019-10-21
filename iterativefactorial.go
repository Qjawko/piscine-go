package piscine

func IterativeFactorial(nb int) int {
	if nb > 65535 {
		return 0
	}

	if nb < 0 {
		return 0
	}

	if nb == 1 || nb == 0 {
		return 1
	}

	result := 1
	for i := nb; i >= 1; i-- {
		result *= i
	}

	return result
}
