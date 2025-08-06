package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age  int   `xml:"age"`
	City string `xml:"city"`
	Email string `xml:"email,attr"`
}

func main() {
	person := Person{
		Name:  "Alice",
		Age:   30,
		City:  "Wonderland",
		Email: "alice@example.com",
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	
	xmlString := xml.Header + string(xmlData)
	fmt.Println(xmlString)
	xmlDataIndented, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to XML with indent:", err)
		return
	}
	
	fmt.Println(string(xmlDataIndented))

	xmlRaw := `<person email="alice@gmail.com"><name>Alice</name><age>30</age><city>Wonderland</city></person>`

	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}
	
	fmt.Printf("Unmarshaled XML: %+v\n", personxml)


}
