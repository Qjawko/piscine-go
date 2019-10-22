package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 || n == 0 {
		z01.PrintRune('0')
	}

	digits := IntToDigits(n)
	sortedDigits := SortIntegerTable(digits)
	PrintIntTable(ConvertArrIntToInt32(sortedDigits))
}

func ConvertArrIntToInt32(arr []int) []int32 {
	var result []int32
	for _, n := range arr {
		result = append(result, int32(n))
	}
	return result
}

func IntToDigits(n int) []int {
	var digits []int
	for n > 0 {
		digit := n % 10
		n /= 10

		digits = append(digits, digit)
	}
	return digits
}

func SortIntegerTableInt(table []int) []int {
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

func ArrLen(runes []int) int {
	count := 0
	for range runes {
		count++
	}

	return count
}
