package piscine

func ConcatParams(args []string) string {
	var result string = args[0]
	for _, s := range args[1:] {
		result += "\n" + s
	}

	return result
}
