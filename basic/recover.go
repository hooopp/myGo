package main

import (
	"fmt"
)

func main() {
	process()
	fmt.Println("Process completed successfully")
}

func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("Starting process...")
	panic("Something went wrong!")
	fmt.Println("Ending process...")
}