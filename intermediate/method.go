package main

import (
	"fmt"
)

func main() {
	rect := Rectangle{length: 5, width: 3}
	area := rect.Area()
	fmt.Println("Area of rectangle:", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area after scaling:", area)

	num1 := MyInt(-5)
	num2 := MyInt(9)
	fmt.Println("Is num1 positive?", num1.IsPositive())
	fmt.Println("Is num2 positive?", num2.IsPositive())
	fmt.Println(num1.welcomeMessage())
	fmt.Println(num2.welcomeMessage())

	s := Shape{
		Rectangle: Rectangle{length: 4, width: 2},
	}
	fmt.Println("Shape area:", s.Area())
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r. width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

func (m MyInt) IsPositive() bool{
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to the world of MyInt!"
}

type Shape struct {
	Rectangle
}