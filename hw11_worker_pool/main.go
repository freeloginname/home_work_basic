package main

import (
	"fmt"
	"sync"
)

func CounterStarter(number int) int {
	var wg sync.WaitGroup
	mx := sync.Mutex{}
	wg.Add(number)
	// для работы через atomic: var counter int64
	var counter int
	for currentGoroutine := 0; currentGoroutine < number; currentGoroutine++ {
		go func(currentGoroutine int) {
			mx.Lock()
			counter++
			// для работы через atomic: atomic.AddInt64(&counter, 1)
			mx.Unlock()
			fmt.Printf("%d goroutine finished with counter %d \n", currentGoroutine, counter)
			wg.Done()
		}(currentGoroutine)
	}
	wg.Wait()
	return counter
}

func main() {
	counter := CounterStarter(10)
	fmt.Printf("Counter: %d \n", counter)
}
