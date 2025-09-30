package main

import (
	"math"
	"testing"
)

func TestConveer(t *testing.T) {
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

	var iteration float64 = 0
	for v := range ch2 {
		if v != math.Pow(iteration, 3) {
			t.Errorf("unexpected number: %.f", v)
		}
		iteration++
	}
}
