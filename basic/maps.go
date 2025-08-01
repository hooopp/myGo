package main

import (
	"fmt"
	"maps"
)

func main() {
	mapVariable := make(map[string]int)

	mapVariable["key1"] = 9
	mapVariable["key2"] = 18
	fmt.Println(mapVariable)

	clear(mapVariable)
	fmt.Println("After clearing:", mapVariable)

	value, unknown := mapVariable["key1"]
	fmt.Println("Value of key1:", value, "Exists:", unknown)

	myMap := map[string]int{
		"key1": 10,
		"key2": 20,
	}

	myMap["newKey"] = 100

	value, exists := myMap["key1"]
	fmt.Println("Value of key1 in myMap:", value, "Exists:", exists)
	value, exists = myMap["key3"]
	fmt.Println("Value of key3 in myMap:", value, "Exists:", exists)

	if maps.Equal(myMap, map[string]int{"key1": 10, "key2": 20}) {
		fmt.Println("myMap is equal to the given map")
	}

	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("myMap4 is nil")
	} else {
		fmt.Println("myMap4 is not nil")
	}

	myMap4 = make(map[string]string)
	myMap5 := map[string]string{}
	fmt.Println("myMap4 and myMap5 are both initialized but empty:", myMap4, myMap5)


}