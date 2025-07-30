package main

import (
	"fmt"
)

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	fruits := [4]string{"apple", "banana", "orange", "grapes"}
	fmt.Println("Fruits array: ",fruits)

	originalArray := [5]int{1, 2, 3, 4, 5}
	copiedArray := originalArray
	copiedArray[0] = 100
	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", copiedArray)

	a, _ := someFunction()
	fmt.Println("Values from someFunction:", a)

	for _, value := range fruits {
		fmt.Println("Fruit:", value)
	}

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{4, 5, 6}

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	originalArray2 := [3]int{1, 2, 3}
	var copiedArray2 *[3]int = &originalArray2
	copiedArray2[0] = 100
	fmt.Println("Original Array2:", originalArray2)
	fmt.Println("Copied Array2:", *copiedArray2)
}

func someFunction() (int, int) {
	return 1, 2
}