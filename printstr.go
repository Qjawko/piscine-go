package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
}
