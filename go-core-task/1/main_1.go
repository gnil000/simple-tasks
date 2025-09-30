package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func main() {
	var intNum10, intNum8, intNum16 int = 42, 052, 0x2A

	fmt.Println(intNum10, intNum8, intNum16)

	var float64Num float64 = 53292.523
	fmt.Println(float64Num)

	var stringSome string = "hello, world"
	fmt.Println(stringSome)

	var boolOp bool = true
	fmt.Println(boolOp)

	var complex64Num complex64 = complex64(6)
	fmt.Println(complex64Num)

	fmt.Println(checkType(intNum10))
	fmt.Println(checkType(intNum8))
	fmt.Println(checkType(intNum16))
	fmt.Println(checkType(float64Num))
	fmt.Println(checkType(stringSome))
	fmt.Println(checkType(boolOp))
	fmt.Println(checkType(complex64Num))

	combined := combineVariablesToString(intNum10, intNum8, intNum16, float64Num, stringSome, boolOp, complex64Num)
	fmt.Println(combined)

	var runeOfSlice []rune = convertStringToSliceOfRune(combined)
	hashed := hashRunes(runeOfSlice)
	fmt.Println(hashed)
}

func checkType(something any) string {
	return reflect.TypeOf(something).String()
}

func combineVariablesToString(intNum10, intNum8, intNum16 int, float64Num float64, stringSome string, boolOp bool, complex64Num complex64) string {
	return fmt.Sprintf("%d %d %d %.2f %s %t %v", intNum10, intNum8, intNum16, float64Num, stringSome, boolOp, complex64Num)
}

func convertStringToSliceOfRune(text string) []rune {
	return []rune(text)
}

func hashRunes(something []rune) string {
	salt := []rune("go-2024")
	mid := len(something) / 2
	somethingSalt := append(something[:mid], append(salt, something[mid:]...)...)
	hash := sha256.Sum256([]byte(string(somethingSalt)))
	return hex.EncodeToString(hash[:])
}
