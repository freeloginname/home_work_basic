package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	sensor_uptime := time.Second * 55
	inputCh := make(chan int64)
	outCh := make(chan int64)
	quitCh := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), sensor_uptime)
	defer func() {
		fmt.Println("Exiting program")
		cancel()
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(inputCh)
				return
			default:
				myRandom, _ := rand.Int(rand.Reader, big.NewInt(64800))
				inputData := myRandom.Int64()
				inputCh <- inputData
			}
		}
	}()

	go func() {
		var sum int64
		var counter int64
		// Если нужно проверять открыт ли канал:
		// data, ok := <-inputCh
		for data := range inputCh {
			fmt.Printf("Getting %d from input channel \n", data)
			sum += data
			counter++
			if counter == 10 {
				fmt.Printf("Sending %d  to output channel \n", sum/counter)
				outCh <- sum / counter
				counter = 0
				sum = 0
			}
		}
		fmt.Println("Input channel closed. Closing output channel.")
		close(outCh)
	}()

	go func() {
		for data := range outCh {
			// data, ok := <-outCh
			// if !ok {
			// 	fmt.Println("Output channel closed. Exiting.")
			// 	close(quitCh)
			// 	return
			// }
			fmt.Printf("Arithmetic mean: %d \n", data)
		}
		fmt.Println("Output channel closed. Exiting.")
		close(quitCh)
	}()
	<-quitCh
}
