package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	go writeToChanel(ch1)
	go writeToChanel(ch2)
	go writeToChanel(ch3)

	for v := range mergeChanels(ch1, ch2, ch3) {
		fmt.Println(v)
	}
}

func mergeChanels(chanels ...<-chan int) <-chan int {
	merged := make(chan int)
	go func() {
		defer close(merged)
		var wg sync.WaitGroup
		wg.Add(len(chanels))
		for _, ch := range chanels {
			go func(local <-chan int) {
				for v := range local {
					merged <- v
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
	}()
	return merged
}

func writeToChanel(ch chan int) {
	defer close(ch)
	for range 10 {
		ch <- rand.Intn(10)
	}
}
