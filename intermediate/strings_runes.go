package main

import (
	"fmt"
	"unicode/utf8"
)
func main() {
	message := "Hello,\nWorld!"
	rawMessage := `Hello,\nWorld!`
	fmt.Println(message)
	fmt.Println(rawMessage)
	fmt.Println("Length of message:", len(message))
	fmt.Println("Length of rawMessage:", len(rawMessage))
	fmt.Println("Rune count: ", utf8.RuneCountInString(message))

	var ch string = "is a"
	var chl string = "is 世"
	chs := '世'
	fmt.Printf("type: %T\n", chs)
	exp := "世界"
	for i,v := range exp {
		fmt.Printf("Character %d: %c with type %T\n", i, v, v)
	}

	fmt.Printf("Character: %v \n", chs) // chs is a rune
	fmt.Printf("Character: %v \n", ch)
	fmt.Printf("Character: %v \n", chl)
	fmt.Println("length of ch:", len(ch))
	fmt.Println("length of chl:", len(chl))
	fmt.Println("Rune count in ch:", utf8.RuneCountInString(ch))
	fmt.Println("Rune count in message:", utf8.RuneCountInString(chl))

}