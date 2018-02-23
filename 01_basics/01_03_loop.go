package main

import "fmt"

func main() {
	// Go only has one type of loop, the 'for' loop. It can be used  few different ways

	// Repeats the command within the 4 loop 10 times
	for i := 0; i < 10; i++ {
		// When the loop starts i will have the value of 0 (i:=0)
		fmt.Println(i + 1)
		// At the end of each loop, i will be increemented by 1 (i++)
	}
	// The loop will continue until i is no longer less than 10 (i < 10)

	// All arguments except the condiion can be skipped, this allows for the values to be changed in the loop code
	// The loop variable must be declared before the loop
	// This format can also be considered the same as other languages 'while' loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Repeats the print command for each element in the 'a' slice
	var a = []string{"a", "b"}

	for i, x := range a {
		fmt.Println(i, x)
	}

	// The 'break' statement can be used to stop the loop at any time (usually with some condition being met)
	for { // This loop will continue indefinately unless something is done to exit the loop
		fmt.Println("I should only appear once")
		break
	}

	// The 'continue' statement can be used to skip the rest of the code in the loop and continue on the next iteration
	for i := 0; i < 10; i++ {
		fmt.Println("See me")
		continue
		fmt.Println("You shouldn't see me")
	}

}
