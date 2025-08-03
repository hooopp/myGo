package main

import (
	"fmt"
)

func main() {
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "New York",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{
			phoneNumber: "123-456-7890",
			homeAddress: "123 Main St, New York, NY",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p2.address.city = "Los Angeles"
	p2.address.country = "USA"

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)

	user := struct {
		username string
		email string
	}{
		username: "johndoe",
		email:    "johndoe@example.com",
	}

	fmt.Println("Username:", user.username)
	fmt.Println("Email:", user.email)

	fmt.Println(p1.fullName())
	p1.incrementAge()
	fmt.Println("After incrementing age:", p1.fullName())

}

type Person struct {
	firstName string
	lastName  string
	age       int
	address Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	phoneNumber string
	homeAddress string
}

type Address struct {
	city string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName + " age is " + fmt.Sprint(p.age)
}

func (p *Person) incrementAge() {
	p.age++
}