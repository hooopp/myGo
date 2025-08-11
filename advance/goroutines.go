package main

import (
	"fmt"
	"time"
)

func main() {
	var err error
	fmt.Println("Starting Goroutine...")
	go sayHello()
	fmt.Print("Hello from Main!")

	go func(){
		err = doWork()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second) // Wait for goroutines to finish

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine!")
}

func printNumbers(){
	for i:=1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Printf("%c\n", letter)
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occurred during work")
}
