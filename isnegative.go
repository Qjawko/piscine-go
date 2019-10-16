package piscine

import (
	"github.com/01-edu/z01"
)

//IsNegative prints is negative number or not
func IsNegative(nb int) {
	if nb > 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune(10)
}
