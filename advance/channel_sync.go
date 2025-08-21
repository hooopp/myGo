package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Done...")

	//multiple working
	nWorking := 5
	fmt.Printf("Starting %d workers...\n", nWorking)
	for i:=0; i<nWorking; i++ {
		go func(i int) {
			fmt.Printf("Working %d...\n", i)
			time.Sleep(2 * time.Second)
			done <- struct{}{}
		}(i)
	}

	for i:=0; i<nWorking; i++ {
		<-done
	}
	fmt.Print("All done...")
}
