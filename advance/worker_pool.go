package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, result chan <- int ){
	for task := range tasks{
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulate work
		result <- task * 2
	}
}

func main() {
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	for range numJobs{
		result := <- results
		fmt.Println("Result:", result)
	}

}
