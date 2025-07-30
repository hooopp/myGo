package main

import "fmt"

func main() {
	var a,b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)
	
	result = a % b
	fmt.Println("Remainder:", result)

	const p1 float64 = 22 / 7
	fmt.Println(p1)

	const p2 float64 = 22 / 7.0
	fmt.Println(p2)

	const p3 float64 = 22.0 / 7
	fmt.Println(p3)

	var maxInt int64 = 9223372036854775807
	fmt.Println("Max Int64:", maxInt)

	maxInt += 1
	fmt.Println("Max Int64 + 1:", maxInt)

}