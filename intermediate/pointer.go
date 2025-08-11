package main

import (
	"fmt"
)

func main() {
	var ptr *int
	var a int = 10
	ptr = &a
	fmt.Println(a)
	fmt.Println(&a)
	modifyValue(&a)
	fmt.Println("Pointer value:", a)
	

	fmt.Println(*ptr)
	modifyValue(ptr)
	fmt.Println("After modification:", *ptr)
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Pointer value:", &ptr)
	fmt.Println("Value of a after modification:", a)
	
}

func modifyValue(input *int) {
	*input++
}