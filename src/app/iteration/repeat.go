package iteration

func Repeat(character string, repeatedCount int) string {
	var repeated string = ""

	for range repeatedCount {
		repeated += character
	}
	return repeated
}
