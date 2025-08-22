package main

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, number int) string {
	select {
	case <- ctx.Done():
		return "Operation canceled"
	default:
		if number%2 == 0 {
			return fmt.Sprintf("%d is even", number)
		}else {
			return fmt.Sprintf("%d is odd", number)
		}
	}
}

func main() {
	// todoContext := context.TODO()

	// ctx := context.WithValue(todoContext, "name", "John")
	// fmt.Println(ctx)
	// fmt.Println(ctx.Value("name"))

	// ctx1 := context.WithValue(todoContext, "city", "New York")
	// fmt.Println(ctx1)
	// fmt.Println(ctx1.Value("city"))

	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // limit time out for context 2 seconds
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result from timeout context:", result)

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result after timeout:", result)

	result = checkEvenOdd(ctx, 20)
	fmt.Println("Result after timeout:", result)

}
