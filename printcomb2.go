package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := 0; i < 99; j++ {
			if i < j {
				if i == 98 && j == 99 {
					z01.PrintRune((rune)(i / 10))
					z01.PrintRune((rune)(i % 10))

					z01.PrintRune(' ')

					z01.PrintRune((rune)(j / 10))
					z01.PrintRune((rune)(j % 10))

					z01.PrintRune(10)
				} else {
					z01.PrintRune((rune)(i / 10))
					z01.PrintRune((rune)(i % 10))

					z01.PrintRune(' ')

					z01.PrintRune((rune)(j / 10))
					z01.PrintRune((rune)(j % 10))
				}
			}
		}
	}
}
