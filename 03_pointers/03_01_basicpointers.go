package main

import "fmt"

// A pointer is used to use hold the memory address of a variable

// Note: Unlike in the C programming lanaguage, pointer arithmetic is not allowed

/*
By default go will use for pass-by-value for all "simple" variables. This means
that a copy of the variable is passed to the function, not the original variable:
   * string
   * int
   * float
   * bool
   * struct

For reference variable types, go will use pass-by-reference so no pointers
are needed:
   * array
   * slice

To reference a variables memory adddress use '&'.

To reference the value of the variable from its address use '*'.

When used with structs, it would seem like a pointer is needed  such as:
(*p).field

Go however allows a shortcut to be made by permitting access without the pointer:
p.field

*/

func main() {
	s := "I am a string"

	// To access the memory address of the string variable 's' use:
	var sp = &s

	/*
		  The following will print the memory address of the variable in hexadecimal
			format:
	*/
	fmt.Printf("%s (%v)\n", s, sp)

	// This is passing the string to the function by-value
	dontChangeMe(s)
	fmt.Printf("%s (%v)\n", s, sp)

	// This is passing the string to the function by-reference
	changeMe(sp)
	fmt.Printf("%s (%v)\n", s, sp)

}

func dontChangeMe(s string) {
	s = "I should not change"
}

func changeMe(sp *string) {
	/*
	  To access the underlying value from a variables address use the '*' operator
	*/
	*sp = "I am now the same string with a different value"
}
