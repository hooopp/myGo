package main

import (
	"fmt"
)

func main() {
	message := "Hello, World!"
	for i,v := range message {
		fmt.Printf("Character at index %d is '%c'\n", i, v)
	}
}