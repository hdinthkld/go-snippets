package main

import "fmt"

// Global variables defined here are accessible by all parts of the program

// Declares a variable called 'n' of type 'int' (a whole number) with default zero value
var n int

// Declares a variable named 's' with an inferred type of 'string'
var s = "I am of type String"

func main() {
	// Declares a local variable of type float with default zero value
	f := 1.2

	// Declares a slice variable of type string with two elements
	var a = []string{"a", "b"}

	fmt.Println(n, f, s)
	fmt.Println(a)

}
