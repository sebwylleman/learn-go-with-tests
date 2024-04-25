package iteration

// Repeat takes a character and returns that character repeated 5 times
func Repeat(character string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += character
	}
	return result
}
