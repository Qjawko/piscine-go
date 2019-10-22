package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 || n == 0 {
		z01.PrintRune('0')
	}

	if n > 2147483647 {
		return
	}


	digits := IntToDigits(int32(n))
	sortedDigits := SortIntegerTableInt32(digits)
	PrintIntTable(sortedDigits)
}

func IntToDigits(n int32) []int32 {
	var digits []int32
	for n > 0 {
		digit := n % 10
		n /= 10

		digits = append(digits, digit)
	}
	return digits
}

func SortIntegerTableInt32(table []int32) []int32 {
	for i := 1; i < ArrLen(table); i++ {
		key := table[i]
		j := i - 1

		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			j = j - 1
		}
		table[j+1] = key
	}
	return table
}

func PrintIntTable(digits []int32) {
	for i := '0'; i <= '9'; i++ {
		for _, digit := range digits {
			if i == digit+48 {
				z01.PrintRune(i)
			}
		}
	}
}

func SortRuneTable(table []rune) []rune {
	for i := 1; i < ArrLen(table); i++ {
		key := table[i]
		j := i - 1

		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			j = j - 1
		}
		table[j+1] = key
	}
	return table
}

func ArrLen(runes []int32) int {
	count := 0
	for range runes {
		count++
	}

	return count
}
