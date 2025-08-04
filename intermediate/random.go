package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(101)) // 0 to 100
	fmt.Println(rand.Intn(100) + 1) // 1 to 100

	val := rand.New(rand.NewSource(42))
	fmt.Println(val.Intn(101))

	val = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(val.Intn(101)) // Random number based on current time
}

