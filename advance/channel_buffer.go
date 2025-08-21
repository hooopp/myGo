package main

import (
	"fmt"
)

func main() {
	sender := make(chan int, 2)
	sender <- 1
	sender <- 2
	receiver := <- sender
	fmt.Println(receiver) // This will block until a receiver is ready
	receiver = <- sender
	fmt.Println(receiver)
}
