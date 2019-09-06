package iteration

func Repeat(char string, numTimes int) string{
	var repeated string
	for i:= 0; i < numTimes; i++ {
		repeated += char
	}
	return repeated
}
