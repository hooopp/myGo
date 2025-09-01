// note:
// mutex use is needed to prevent race conditions when multiple goroutines access

package main

import (
	"fmt"
	"sync"
)
type counter struct {
	mu sync.Mutex
	count  int
}

func (c * counter) increment(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}


func main() {
	var wg sync.WaitGroup
	counter := &counter{}
	numGoroutines := 1000

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter Value:", counter.getValue())
	counter.count = 0
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.count++
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter Value:", counter.getValue())
}
