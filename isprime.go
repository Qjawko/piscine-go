package piscine

func IsPrime(n int) bool {
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
