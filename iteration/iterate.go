package iteration

// Repeat - this is a function that does blah
func Repeat(char string, amt int) string {

	var repeated string
	for i := 0; i < amt; i++ {
		repeated += char
	}
	return repeated
}

// func main() {
// 	fmt.Println(strings.Index("   .     hello     I am Kim     ", "h"))
// }
