package main

import (
	"fmt"
	"sync"
	"time"
)

// func worker(id int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Printf("Worker %d started\n", id)
// 	time.Sleep(time.Second) // Simulate work
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() //blocking mechanism
// 	fmt.Println("All workers finished")
// }

// func worker(id int, results chan<-int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting. \n", id)
// 	time.Sleep(time.Second)
// 	results <- id * 2
// 	fmt.Printf("WorkerID %d finished. \n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 3
// 	results := make(chan int, numJobs)

// 	wg.Add(numJobs)

// 	for i := range numWorkers {
// 		go worker(i, results, &wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Printf("Result: %d\n", result)
// 	}

// }

type Worker struct {
	ID int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started task: %s\n", w.ID, w.Task)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d finished task: %s\n", w.ID, w.Task)
}

func main(){
	var wg sync.WaitGroup
	tasks := []string {"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i+1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}
	wg.Wait()
	fmt.Println("All workers finished")
}