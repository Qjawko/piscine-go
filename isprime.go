package piscine

//https://en.wikipedia.org/wiki/Primality_test
func IsPrime(n int) bool {
	if n < 0 {
		return false
	}

	if n <= 3 {
		return n > 1
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}

	return true
}
