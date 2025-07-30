package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	sum := 0
	for {
		sum += 10
		fmt.Println("Current Sum:", sum)
		if sum >= 50 {
			break
		}
	}

	num := 1
	for num <= 10 {
		if num %2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number:", num)
		num++
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1
	fmt.Println("Guess the number between 1 and 100")
	fmt.Println("I have chosen a number, try to guess it!")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("Congratulations! You've guessed the number:", target)
			break
		}else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}