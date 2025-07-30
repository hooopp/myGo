package main

import (
	"fmt"
)

func main() {
	fruit := "apple"

	switch fruit {
		case "apple":
			fmt.Println("This is an apple.")
		case "banana":
			fmt.Println("This is a banana.")
		default: 
			fmt.Println("Unknown fruit.")
	}

	day := "Monday"
	switch day {
		case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("It's a weekday.")
		case "Saturday", "Sunday":
			fmt.Println("It's the weekend.")
		default:
			fmt.Println("Unknown day.")
	}

	number := 23

	switch {
		case number < 10:
			fmt.Println("Number is less than 10.")
		case number >= 10 && number < 20:
			fmt.Println("Number is between 10 and 20.")
		default:
			fmt.Println("Number is 20 or greater.")
	}

	checkType(42)
	checkType(3.14)
	checkType("Hello, Go!")

}

func checkType (x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown type")
	}
}