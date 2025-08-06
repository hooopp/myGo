package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command: ", os.Args[0])
	for i, arg := range os.Args {
		fmt.Println("Argument: ", i, ":", arg)
	}

	var name string
	var age int

	flag.StringVar(&name, "name", "John Doe", "description: Your name")
	flag.IntVar(&age, "age", 30, "description: Your age")

	flag.Parse()

	fmt.Println("Parsed Arguments:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

}
