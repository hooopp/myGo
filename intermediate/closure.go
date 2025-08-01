package main

import (
	"fmt"
)

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println("Subtracting 1:", subtracter(1))
	fmt.Println("Subtracting 2:", subtracter(2))
	fmt.Println("Subtracting 3:", subtracter(3))

}

func adder() func () int {
	i := 0
	fmt.Println("previous value:", i)
	return func() int {
		i++
		fmt.Println("current value:", i)
		return i
	}
}