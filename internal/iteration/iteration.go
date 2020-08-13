package iteration

import "fmt"

// Repeat : Repeats a letter
func Repeat(input string, repeats int) string {
	var repeated string
	for i := 0; i < repeats; i++ {
		repeated += input
	}
	return repeated
}

// ExampleRepeat : an example of how to use the Add function
func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output : aaaaa
}
