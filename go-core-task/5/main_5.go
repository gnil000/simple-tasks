package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(hasIntersections(a, b))
}

func hasIntersections(slice1, slice2 []int) (bool, []int) {
	intersection := make([]int, 0)

	setFirstSlice := make(map[int]struct{})
	for _, v := range slice1 {
		setFirstSlice[v] = struct{}{}
	}
	existsElements := make(map[int]struct{})
	for _, v := range slice2 {
		if _, ok := setFirstSlice[v]; ok {
			if _, exists := existsElements[v]; !exists {
				intersection = append(intersection, v)
				existsElements[v] = struct{}{}
			}
		}
	}

	return len(intersection) > 0, intersection
}
