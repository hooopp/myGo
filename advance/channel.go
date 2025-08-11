package main

import (
	"fmt"
	"time"
)

func main() {
	greeting := make(chan string)

	go func(){
		greeting <- "Hello, World!"
		time.Sleep(1 * time.Second) // Simulate some work
		greeting <- "Goroutine finished sending greeting"
	}()

	receiver := <-greeting
	fmt.Println(receiver)
	time.Sleep(2 * time.Second) // Wait for goroutine to finish
	receiver = <-greeting
	fmt.Println(receiver)
	fmt.Println("End of program")

	go func() {
		greeting <- "Hello"
		greeting <- "world"
	}()

	go func() {
		receiver := <-greeting
		fmt.Println(receiver)
		receiver = <-greeting
		fmt.Println(receiver)
	}()

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println("All greetings received")


}