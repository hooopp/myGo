package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting app")
	timer := time.NewTimer(2 * time.Second) // non blocking
	fmt.Println("Waiting for timer...")
	<-timer.C // blocking wait for timer
	fmt.Println("Timer expired")
}
  