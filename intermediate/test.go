package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ptr *int
	ptr = &a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of ptr:", *ptr)
	modifyValue(ptr)
	fmt.Println("After modification, value pointed by ptr:", *ptr)
	fmt.Println("Value of a after modification:", a)

}

func modifyValue(input *int) {
	*input++
}