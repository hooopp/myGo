package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()
	
	select { // blocking until any case receive for channel
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
	}
	select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
		default: // using for non blocking
			fmt.Println("No message received")
	}
	//limit time out
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
	// handle channel close

	go func(){
		ch1 <- 1
		close(ch1)
	}()
	for {
		select {
			case msg, ok := <-ch1:
				fmt.Println("status:", ok)
				if !ok {
					fmt.Println("Channel closed")
					return 
				}
				fmt.Println("Received from ch1:", msg)
		}
	}

}
