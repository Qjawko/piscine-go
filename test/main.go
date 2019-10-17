package main

import (
	piscine ".."
	"fmt"
)

func main() {
	str := "Hello World!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
