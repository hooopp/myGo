package main

import (
	"fmt"
)
func main() {
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Factorial of 10:", factorial(10))

	fmt.Println("Sum of digits in 12:", sumOfDigits(12))
	fmt.Println("Sum of digits in 123:", sumOfDigits(123))
}

func factorial(n int) int {
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}