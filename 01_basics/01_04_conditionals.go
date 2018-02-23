package main

import "fmt"

func main() {

	// Go's conditional statement is 'if'.  Its most basic form is

	x := 1
	if x == 1 {
		// This statement will only be executed if the above statement is true
		fmt.Println("I agree that the statement is true")
	}

	// Multiple matches can be done with 'else', each condition checked in turn until one is true
	x = 4
	if x == 1 {
		fmt.Println("One")
	} else if x == 2 {
		fmt.Println("Two")
	} else if x == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Nope, haven't got a clue what to do with", x)
	}

	// A simple statement can also be run before the condition is checked
	if x = 1; x > 0 {
		fmt.Println("Well at least we have something")
	}

	// To run code based on a certain set of values 'switch' can be used
	switch z := 4; z {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Nope, haven't got a clue what to do with", z)
	}
}
