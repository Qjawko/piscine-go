package piscine

func ConcatParams(args []string) string {
	var result string
	for _, s := range args {
		result += s
	}

	return result
}
