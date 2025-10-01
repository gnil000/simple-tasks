package main

import (
	"fmt"
	"math"
)

func main() {
	ch1 := make(chan uint8)
	var size uint8 = 10
	go func() {
		for i := range size {
			ch1 <- i
		}
		close(ch1)
	}()

	ch2 := make(chan float64)
	go conveer(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}
}

func conveer(uintChan <-chan uint8, floatChan chan<- float64) {
	for v := range uintChan {
		floatChan <- math.Pow(float64(v), 3)
	}
	close(floatChan)
}
