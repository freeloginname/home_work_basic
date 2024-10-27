package hw10motionsensor

import (
	"fmt"
	"sync"
	"time"
)

func Reader() {
	var inputData int
	fmt.Scanf("%d \n", &inputData)
}

func main() {
	//var wg sync.WaitGroup
	var inputData int
	mx := sync.Mutex{}
	inputCh := make(chan int)
	outCh := make(chan int)
	//wg.Add(1)
	currentTime := time.Now()
	fmt.Printf("Current time: %s \n", currentTime)
	timeout := currentTime.Add(10 * time.Second)

	go func() {
		for currentTime.Before(timeout) {
			fmt.Scanf("%d \n", &inputData)
			fmt.Printf("Sending %d to input channel \n", inputData)
			inputCh <- inputData
		}
	}()

	go func() {
		sum := 0
		counter := 0
		for {
			//var data int
			data := <-inputCh
			fmt.Printf("Getting %d from input channel \n", data)
			mx.Lock()
			sum = sum + data
			counter++
			fmt.Printf("Sending %d  to output channel \n", sum/counter)
			outCh <- sum / counter
			mx.Unlock()

		}
	}()

	go func() {
		for {
			fmt.Printf("Printing middle")
			fmt.Println(<-outCh)
		}
	}()
	time.Sleep(time.Minute)
	//wg.Wait()
	defer fmt.Printf("Result: %d", <-outCh)
}
