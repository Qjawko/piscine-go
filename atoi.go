package piscine

import (
	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if StrLen(s) < 1 {
			return 0
		}
	}

	nm := 0

	for _, ch := range s {
		z01.PrintRune(ch)
		if !containsIn0to9(ch) {
			return 0
		}
		nm = nm*10 + charToInt(ch)
	}

	if s0[0] == '-' {
		nm *= -1
	}
	return nm
}
