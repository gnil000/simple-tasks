package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := createRandomSlice()
	fmt.Println(originalSlice)

	fmt.Println(sliceExample(originalSlice))

	fmt.Println(addElements(originalSlice, 7))

	copyS := copySlice(originalSlice)
	copyS[0] = 111
	fmt.Println(copyS, originalSlice) // изменения не повлияли

	sliceRem := createRandomSlice()
	fmt.Println(sliceRem, removeElement(sliceRem, 0))
}

func createRandomSlice() []int {
	slice := make([]int, 0)
	for range 10 {
		slice = append(slice, rand.Intn(10))
	}
	return slice
}

func sliceExample(slice []int) []int {
	resultSlice := make([]int, 0, cap(slice))
	for _, v := range slice {
		if v%2 == 0 {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	copyS := make([]int, len(slice))
	copy(copyS, slice)
	return copyS
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
