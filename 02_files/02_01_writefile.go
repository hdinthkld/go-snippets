package main

import (
	"io/ioutil"
)

func main() {

	// The simplest example is to write out a string directly as a series of bytes to a file
	fn := "temp/test.txt"                        // The file to save to
	d := "I am a test string"                    // The data to write to the file
	err := ioutil.WriteFile(fn, []byte(d), 0777) // Writes out the data to the file

	// Exit the program in the event of any error
	if err != nil {
		panic(err)
	}

}
