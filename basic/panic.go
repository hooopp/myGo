package main

import (
	"fmt"
)

func main() {
	// process(5)
	process(-1)
	// process(10) 
}

func process(input int){
	defer fmt.Println("Deffered!!!")
	if input < 0 {
		panic("Negative input not allowed")
		defer fmt.Println("Deffered 2") // This will not execute
	}
	fmt.Println("Processing input:", input)
}