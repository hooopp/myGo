package main

import (
	"fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    data := map[string]interface{}{
		"name": "Jane",
		"age":  25,
	}
    fmt.Println("Welcome to", data["name"], "who is", data["age"], "years old!")
}
