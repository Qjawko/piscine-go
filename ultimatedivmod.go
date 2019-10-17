package piscine

func UltimateDivMod(a, b *int) {
	var buffer int
	buffer = *a / *b
	*b = *a % *b
	*a = buffer
}
