package iteration

const repeatCount = 5

func Repeat(character string) string {
	result := ""

	for i := 0; i < repeatCount; i++ {
		result += character
	}

	return result

}
