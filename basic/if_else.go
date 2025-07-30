package main

import "fmt"

func main() {
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	temperature := 25
	if temperature >= 30{
		fmt.Println("It's a hot day.")
	}else {
		fmt.Println("It's a cool day.")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	}else if score >= 80 {
		fmt.Println("Grade B")
	}else if score >= 70 {
		fmt.Println("Grade C")
	}else{
		fmt.Println("Grade D")
	}

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if num % 2 == 0 {
		if num % 3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3")
		} else {
			fmt.Println("Number is divisible by 2 but not by 3")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}
}