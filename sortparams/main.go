package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	joinedString := BasicJoin(os.Args[1:])
	sortedRunes := SortRuneTable([]rune(joinedString))
	PrintRuneArray(sortedRunes)
}

func SortRuneTable(table []rune) []rune {
	for i := 1; i < RuneArrLen(table); i++ {
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

func PrintRuneArray(runes []rune) {
	for _, ch := range runes {
		_ = z01.PrintRune(ch)
		_ = z01.PrintRune('\n')
	}
}

func BasicJoin(strs []string) string {
	var result string
	for _, str := range strs {
		result += str
	}
	return result
}

func RuneArrLen(arr []rune) int {
	count := 0
	for range arr {
		count++
	}
	return count
}
