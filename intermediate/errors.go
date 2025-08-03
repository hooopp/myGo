package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot take square root of negative number")
	}
	return 1, nil
}


func main() {
	result, err := sqrt(-4)
	if err != nil{
		fmt.Print(err)
		return
	}
	fmt.Println(result)

	result, err1 := sqrt(16)
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	fmt.Println(result)

}