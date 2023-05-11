package iteration

func Repeat(character string, repeatcount int) string {
	var repeated string
	for i := 0; i < repeatcount; i++ {
		repeated += character
	}
	return repeated
}
