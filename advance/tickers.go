package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	i := 0
	for tick := range ticker.C {
		fmt.Println(i, "Tick at", tick)
		i++
	}
}
