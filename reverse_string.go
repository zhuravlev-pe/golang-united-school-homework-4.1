package reverse_string

func reverseSlice(runes []rune) {
	if len(runes) < 2 {
		return
	}
	to := len(runes) / 2

	for i := 0; i < to; i++ {
		with := len(runes) - 1 - i
		runes[i], runes[with] = runes[with], runes[i]
	}
}

func ReverseString(input string) (output string) {
	runes := []rune(input)
	reverseSlice(runes)
	output = string(runes)
	return output
}
