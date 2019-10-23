package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		PrintRuneArray([]rune(s))
		z01.PrintRune('\n')
	}
}

func PrintRuneArray(runes []rune) {
	for _, ch := range runes {
		_ = z01.PrintRune(ch)
	}
}
