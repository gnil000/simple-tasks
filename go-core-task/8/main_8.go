package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomWaitGroup struct {
	counter int
	cond    *sync.Cond
}

func NewWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{cond: sync.NewCond(&sync.Mutex{})}
}

func (wg *CustomWaitGroup) Add(n int) {
	wg.cond.L.Lock()
	wg.counter += n
	wg.cond.L.Unlock()
}

func (wg *CustomWaitGroup) Done() {
	wg.cond.L.Lock()
	wg.counter--
	wg.cond.Signal()
	wg.cond.L.Unlock()
}

func (wg *CustomWaitGroup) Wait() {
	wg.cond.L.Lock()
	for wg.counter > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func work(i int, wg *CustomWaitGroup) {
	fmt.Println(" work ", i, " started.")
	time.Sleep(1 * time.Second)
	fmt.Println(" work ", i, " ended.")
	wg.Done()
}

func main() {
	wg := NewWaitGroup()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()
}
