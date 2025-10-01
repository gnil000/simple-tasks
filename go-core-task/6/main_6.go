package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := randNumsGenerator(10)
	for v := range nums {
		fmt.Println("Num: ", v)
	}
}

func randNumsGenerator(count int) <-chan int {
	nums := make(chan int)
	go func() {
		for range count {
			nums <- rand.Int()
		}
		close(nums)
	}()
	return nums
}
