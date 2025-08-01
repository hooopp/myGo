package main

import (
	"fmt"
	"slices"
)

func main() {
	// var numbers1 []int
	// var numbers2 []int{1,2,3}

	// numbers3 := []int{7,8,9}

	slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice = a[1:4] // slice from index 1 to 3
	fmt.Println("slice:", slice)

	slice = append(slice, 6, 7)
	fmt.Println("slice after append:", slice)

	sliceCopy := make([]int, len(slice))
	fmt.Println("sliceCopy before copy:", sliceCopy)
	copy(sliceCopy, slice)
	fmt.Println("sliceCopy after copy:", sliceCopy)

	// var nilSlice []int

	for i, v := range slice{
		fmt.Println("Index:", i, "Value:", v)
	}

	if slices.Equal(slice, sliceCopy) {
		fmt.Println("Slices are equal")
	}

	slice2D := make([][]int, 3)
	fmt.Println("2D Slice before assignment:", slice2D)
	

}