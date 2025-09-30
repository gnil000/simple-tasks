package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	// result = {"apple", "cherry", "43", "lead", "gno1"}
	fmt.Println(LeftJoinSlices(slice1, slice2))
}

func LeftJoinSlices(slice1, slice2 []string) []string {
	slice := make([]string, 0)
	for _, v1 := range slice1 {
		flag := true
		for _, v2 := range slice2 {
			if v1 == v2 {
				flag = false
				break
			}
		}
		if flag {
			slice = append(slice, v1)
		}
	}
	return slice
}
