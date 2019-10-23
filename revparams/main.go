package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	for i := StringArrLen(os.Args) - 1; i > 0; i-- {
		PrintRuneArray([]rune(os.Args[i]))
		z01.PrintRune('\n')
	}
}

func StringArrLen(arr []string) int {
	count := 0
	for range arr {
		count++
	}
	return count
}

func PrintRuneArray(runes []rune) {
	for _, ch := range runes {
		_ = z01.PrintRune(ch)
	}
}
