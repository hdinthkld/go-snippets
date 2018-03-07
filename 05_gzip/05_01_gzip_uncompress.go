package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	fnIn := "./test.xml.gz" // The source GZIP compressed file
	fnOut := "./test.xml"   // The destination uncompressed file

	// Open the compressed file
	gzInFile, err := os.Open(fnIn)

	if err != nil {
		os.Exit(1)
	}
	defer gzInFile.Close()

	// Create the output file
	gzOutFile, err := os.Create(fnOut)
	if err != nil {
		os.Exit(2)
	}
	defer gzOutFile.Close()

	// Setup the way to read in the compressed data
	zr, err := gzip.NewReader(gzInFile)
	if err != nil {
		os.Exit(3)
	}
	defer zr.Close()

	// Copy the uncompressed data to the output file
	_, err = io.Copy(gzOutFile, zr)

}
