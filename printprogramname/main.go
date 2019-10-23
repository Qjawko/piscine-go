package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	PrintRuneArray([]rune(os.Args[0]))
	_ = z01.PrintRune('\n')
}

func PrintRuneArray(runes []rune) {
	for _, ch := range runes {
		_ = z01.PrintRune(ch)
	}
}
