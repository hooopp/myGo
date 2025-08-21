package main

import "fmt"

func main() {
    // Initialize an empty slice to act as our stack
    stack := []string{}

    // Push elements onto the stack
    stack = append(stack, "apple")
    stack = append(stack, "banana")
    stack = append(stack, "cherry")

    fmt.Println("Stack after pushes:", stack)

    // Peek at the top element without removing it
    top := stack[len(stack)-1]
    fmt.Println("Top element is:", top)

    // Pop the top element
    popped := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    fmt.Println("Popped element:", popped)
    fmt.Println("Stack after pop:", stack)
}