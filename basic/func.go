package main

import (
	"fmt"
)

func main() {

	fmt.Println(add(2,3))

	func() {
		fmt.Println("Anonymous function called")
	}()

	greet := func(){
		fmt.Println("Anonymous function called")
	}
	greet()

	operation := add

	result := operation(3,5)
	fmt.Println("Result of operation:", result)

	result1 := applyOperation(4, 6, add)
	fmt.Println("Result of applyOperation:", result1)

	multiplyByTwo := createMultiplier(2)
	fmt.Println("Multiplying 5 by 2:", multiplyByTwo(5))

}

func add(a , b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}