package main

import "fmt"

func main() {

	for i:=1; i<=5; i++ {
		fmt.Println("Iteration:", i)
	}

	numbers := []int{1,2,3}
	for index, value := range numbers{
		fmt.Println("Index:", index, "Value:", value)
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i:=1; i<=10; i++{
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break
		}
	}

	rows := 5

	for i:=1;i<=rows;i++{
		for j:=1; j<=rows-i;j++{
			fmt.Print(" ")
		}
		for k:=1;k<=2*i-1;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := range 5{
		fmt.Println("Range Loop:", i)
	}
}