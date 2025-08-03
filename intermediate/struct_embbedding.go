package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type Employee struct {
	empployeeInfo person
	empId string
	salary float64
}

func (p person) introduce(){
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, my employee ID is %s and my salary is %.2f.\n", e.empployeeInfo.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		empployeeInfo: person{name: "john", age: 30},
		empId: "E001", 
		salary: 50000.0,
	}
	fmt.Println("Employee Name:", emp.empployeeInfo.name)
	fmt.Println("Employee Age:", emp.empployeeInfo.age)
	fmt.Println("Employee ID:", emp.empId)
	fmt.Println("Employee Salary:", emp.salary)

	emp.introduce()
	emp.empployeeInfo.introduce()
}
