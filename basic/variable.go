package main

import "fmt"

var middleName = "Cane"

func main() {
	var age int
	fmt.Println("age: ", age)
	var name1 string = "john"
	fmt.Println("Hello" + name1)
	var name2 = "john"
	fmt.Println("Hello " + name2)

	count := 10
	fmt.Println("Count:", count)
	lastName := "Smith"
	fmt.Println("Last Name:", lastName)

	fmt.Println("middle Name:", middleName)
	middleName := "Mayor"
	fmt.Println("middle Name:", middleName)

}

func printname(){
	firstName := "Michael"
	fmt.Println(firstName)
}