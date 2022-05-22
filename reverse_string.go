package reverse_string



func ReverseString(input string) (output string) {
	r := []rune(input)
	r1 := []rune("")
	
	for i := len(r)-1; i >= 0; i-- {
		r1 = append(r1, r[i])
	}
	
	return string(r1)
}
