package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	// Notify channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		sig := <- sigs
		fmt.Println("Received signal:", sig)
		fmt.Println("Exiting gracefully...")
		os.Exit(0)
	}()

	fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}

}
