package main

import (
	"fmt"
)

func main() {
	process(2)
	fmt.Println("Main function executed")
}

func process(i int) {
	defer fmt.Println("Deffered i vale", i)
	defer fmt.Println("First deferred function executed")
	defer fmt.Println("Second deferred function executed")
	i++
	fmt.Println("Process function executed")
	fmt.Println("value of i:", i)
}