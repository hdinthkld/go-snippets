// NOTE: Make sure you've run 02_01_writefile.go before running this program
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// This example code reads a file into a variable in one go
	// All the action of opening, reading and closing the file is done in one action
	var fn = "temp/test.txt"    // File to open
	d, e := ioutil.ReadFile(fn) // Opens and reads the data into a variable

	if e != nil {
		panic(e)
	}

	fmt.Println(string(d))

}
