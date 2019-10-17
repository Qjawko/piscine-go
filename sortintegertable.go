package piscine

func SortIntegerTable(table []int) []int {
	for i := 1; i < len(table); i++ {
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
