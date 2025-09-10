package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	t2 := v.Type()

	fmt.Println("Type from TypeOf:", t)
	fmt.Println("Type from Value.Type():", t2)
	
	fmt.Println("Value:", v)

	fmt.Println("Is Integer:", v.Kind() == reflect.Int)
}
