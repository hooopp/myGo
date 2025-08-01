package main

import (
	"fmt"
)
func main() {
	s1 := fmt.Sprint("Hello ", "World! ", 123, 456)
	fmt.Println(s1)
	s2 := fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Println(s2)
	s3 := fmt.Sprintf("Hello %s! %d %d", "World", 123, 456)
	fmt.Println(s3)

	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age) 


}