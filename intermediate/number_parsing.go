package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing number:", err)
		return
	}
	fmt.Println("Parsed number:", num)
	fmt.Println("Parsed number:", num+1)

	// Example of parsing a float
	floatStr := "123.45"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	fmt.Println("Parsed float:", floatNum)

}