package main

import "testing"

func TestMergeChanels(t *testing.T) {
	sizeChanel := 10
	sizeSlice := sizeChanel * 3
	ch1, ch2, ch3 := make(chan int, sizeChanel), make(chan int, sizeChanel), make(chan int, sizeChanel)
	for i := range sizeChanel {
		ch1 <- i
	}
	close(ch1)
	for i := range sizeChanel {
		ch2 <- i
	}
	close(ch2)
	for i := range sizeChanel {
		ch3 <- i
	}
	close(ch3)

	resultCh := mergeChanels(ch1, ch2, ch3)

	resultSlice := make([]int, 0)
	for v := range resultCh {
		resultSlice = append(resultSlice, v)
	}

	if len(resultSlice) != sizeSlice {
		t.Errorf("expected length %d, but got: %d", sizeSlice, len(resultSlice))
	}

	counts := make(map[int]int)
	for _, v := range resultSlice {
		counts[v]++
	}

	for i := range sizeChanel {
		if n, ok := counts[i]; ok {
			if n != 3 {
				t.Errorf("Every num must be in 3 copies")
			}
		}
	}
}
