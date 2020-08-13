package main

import "fmt"

// Add : takes two integers and adds them together
func Add(x, y int) int {
	return x + y
}

// ExampleAdd : an example of how to use the Add function
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output : 6
}
