package iteration

// Repeat takes a character and returns that character repeated n times
func Repeat(character string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += character
	}
	return result
}
