package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	mx := sync.Mutex{}
	inputCh := make(chan int64)
	outCh := make(chan int64)
	// wg.Add(1)

	go func() {
		for {
			myRandom, _ := rand.Int(rand.Reader, big.NewInt(64800))
			inputData := myRandom.Int64()
			inputCh <- inputData
		}
	}()

	go func() {
		var sum int64
		var counter int64
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
