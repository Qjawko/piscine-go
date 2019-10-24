package main

import (
	"fmt"
	"os"
)

func main() {
	isOrder, _ := containsArg(os.Args[1:], "--order", 7, false)
	isInsert, insertArg := containsArg(os.Args[1:], "--insert", 8, true)
	isHelp, _ := containsArg(os.Args[1:], "--help", 6, false)
	if isHelp {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("    This flag inserts the string passed as argument.")

		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("    This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}

	otherArgs := getArgs(os.Args[1:])

	for _, arg := range otherArgs {
		var runes []rune
		if isOrder {
			runes = sortRuneTable([]rune(arg))
		} else {
			runes = []rune(arg)
		}
		fmt.Print(string(runes))
	}

	if isInsert {
		if isOrder {
			fmt.Println(string(sortRuneTable([]rune(insertArg))))
		} else {
			fmt.Println(insertArg)
		}
	}
}

func sortRuneTable(table []rune) []rune {
	for i := 1; i < stringLen(table); i++ {
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

func stringLen(s []rune) int {
	count := 0
	for range s {
		count++
	}

	return count
}

func getArgs(args []string) []string {
	var result []string

	for _, arg := range args {
		if !checkForRune(arg, '-') {
			result = append(result, arg)
		}
	}

	return result
}

func checkForRune(s string, check rune) bool {
	for _, r := range s {
		if r == check {
			return true
		}
	}

	return false
}

//func getInsertValue(s string) []rune {
//	fmt.Println(s)
//	runes := []rune(s)
//	for i, value := range runes {
//		if value == '=' {
//			return
//		}
//	}
//
//	return nil
//}

func containsArg(args []string, argToFind string, argToFindLen int, takeParam bool) (bool, string) {
	for _, arg := range args {

		if checkForRune(arg, '-') {
			if arg[:argToFindLen] == argToFind || arg == argToFind[1:3] {
				if takeParam {
					return true, arg[argToFindLen+1:]
				} else {
					return true, ""
				}
			}
		}
	}

	return false, ""
}
