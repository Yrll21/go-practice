package iteration

func Repeat(char string, repeatitions int) (output string) {
	i := 0
	for i < repeatitions {
		output += char
		i++
	}
	return
}
