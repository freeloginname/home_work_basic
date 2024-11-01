package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func CounterStarter(number int) int64 {
	var wg sync.WaitGroup
	mx := sync.Mutex{}
	wg.Add(number)
	// var counter int
	var counter int64
	for currentGoroutine := 0; currentGoroutine < number; currentGoroutine++ {
		go func(currentGoroutine int) {
			mx.Lock()
			// counter++
			atomic.AddInt64(&counter, 1)
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
