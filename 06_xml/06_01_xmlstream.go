package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Credits
// -------
// Parsing Huge XML Files With Go
// http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/

type Redirect struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Title string   `xml:"title"`
	Redir Redirect `xml:"redirect"`
	Text  string   `xml:"revision>text"`
}

func main() {
	var inputFile string
	inputFile = "./wiki.xml"

	xmlFile, err := os.Open(inputFile)
	if err != nil {
		os.Exit(1)
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	total := 0
	var inElement string

	for {
		t, _ := decoder.Token()

		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			inElement = se.Name.Local
			if inElement == "page" {
				var p Page
				decoder.DecodeElement(&p, &se)
				fmt.Println(p)
			}
		default:
		}
	}
	fmt.Printf("Total articles: %d \n", total)
}
