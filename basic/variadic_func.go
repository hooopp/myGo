package main

import (
	"fmt"
)

func main() {

	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	numbers := []int{4, 5, 6, 7}
	fmt.Println("Sum of 4, 5, 6, 7:", sum(numbers...))

}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}