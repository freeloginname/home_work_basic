package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	rand.NewSource(time.Now().UnixNano())
	mx := sync.Mutex{}
	inputCh := make(chan int)
	outCh := make(chan int)
	// wg.Add(1)

	go func() {
		for {
			inputData := rand.Intn(100)
			inputCh <- inputData
		}
	}()

	go func() {
		sum := 0
		counter := 0
		for {
			// var data int
			data := <-inputCh
			fmt.Printf("Getting %d from input channel \n", data)
			mx.Lock()
			sum += data
			counter++
			mx.Unlock()
			if counter == 10 {
				fmt.Printf("Sending %d  to output channel \n", sum/counter)
				outCh <- sum / counter
				counter = 0
				sum = 0
			}
		}
	}()

	go func() {
		for {
			fmt.Printf("Arithmetic mean: %d \n", <-outCh)
		}
	}()
	time.Sleep(time.Minute)
	// time.Sleep(10 * time.Second)
	// wg.Wait()
}
