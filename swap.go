package piscine

func Swap(a *int, b *int) {
	buffer := *a
	*a = *b
	*b = buffer
}
