package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received from channel:", msg)
	// default:
	// 	fmt.Println("No message received")
	// }
	// select {
	// 	case ch <- 1:
	// 		fmt.Println("Sent to channel")
	// 	default:
	// 		fmt.Println("Channel is not ready")
	// }

	// non blocking
	data := make(chan int)
	quit := make(chan bool)

	go func(){
		for {
			select {
			case d := <-data:
				fmt.Println("Received data:", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
	time.Sleep(time.Second) // wait for goroutine to finish
}
