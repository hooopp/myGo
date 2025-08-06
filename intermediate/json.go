package main

import (
	"encoding/json"
	"fmt"
)

type Adress struct {
	City string `json:"city"`
	State string `json:"state"`
}

type Person struct {
	FirstName string `json:"name"`
	Age      int    `json:"age"`
	Email   string `json:"email,omitempty"`
	Address Adress `json:"address"`
}

type Employee struct {
	FullName string `json:"full_name"`
	EmpID    string `json:"emp_id"`
	Age      int    `json:"age"`
	Address  struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
}

func main() {
	person := Person{
		FirstName: "Alice",
		Address: Adress{
			City:  "Wonderland",
			State: "Fantasy",
		},
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println("JSON data:", string(jsonData))

	jsonDataString := `{"full_name": "jenny Doe","emp_id": "0009","age": 30,"address": {"city": "New York","state": "NY"}}`

	var employee Employee
	err = json.Unmarshal([]byte(jsonDataString), &employee)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Unmarshalled Employee:", employee)
}

