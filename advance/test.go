package main

import (
	"fmt"
)

func main() {
	greeting := make(chan string, 1) // Buffer size 1
	greeting <- "Hello, World!"      // Won't block
	var receive string
	receive = <-greeting
	fmt.Println(receive)
}