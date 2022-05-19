package iteration

func Repeat(str string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += str
	}
	return repeated
}
