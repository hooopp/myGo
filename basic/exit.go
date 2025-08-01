package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting the program...")
	defer fmt.Println("Deferred function executed before exit")

	os.Exit(1)

	fmt.Println("This line will not be executed due to os.Exit")
}