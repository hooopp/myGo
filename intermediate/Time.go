package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0 , time.UTC)
	fmt.Println(specificTime)

	parseTimes, _ := time.Parse("2006-01-02", "2024-07-30")
	parseTimes1, _ := time.Parse("06-01-02", "20-05-01")
	fmt.Println(parseTimes)
	fmt.Println(parseTimes1)

	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("2006-01-02 15:04:05"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())
	fmt.Println("Round Time: ", t.Round(time.Hour))
	
}