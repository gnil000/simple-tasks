package main

import (
	"testing"
	"time"
)

func TestCorrectWay(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(10)
	for range 10 {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
		}()
	}
	wg.Wait()
}

func TestCorrectWay_Reuse(t *testing.T) {
	wg := NewWaitGroup()

	wg.Add(2)
	wg.Done()
	wg.Done()

	wg.Wait()

	wg.Add(1)
	wg.Done()

	wg.Wait()
}
