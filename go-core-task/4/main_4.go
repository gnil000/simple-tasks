package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	// result = {"apple", "cherry", "43", "lead", "gno1"}
	fmt.Println(leftJoinSlices(slice1, slice2))
}

// по умному с мапой
func leftJoinSlices(slice1, slice2 []string) []string {
	slice := make([]string, 0)
	setFromSlice2 := make(map[string]struct{})
	for _, v := range slice2 {
		setFromSlice2[v] = struct{}{}
	}
	for _, v := range slice1 {
		if _, exists := setFromSlice2[v]; !exists {
			slice = append(slice, v)
		}
	}
	return slice
}

// реализация в тупую
// func leftJoinSlices(slice1, slice2 []string) []string {
// 	slice := make([]string, 0)
// 	for _, v1 := range slice1 {
// 		flag := true
// 		for _, v2 := range slice2 {
// 			if v1 == v2 {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			slice = append(slice, v1)
// 		}
// 	}
// 	return slice
// }
