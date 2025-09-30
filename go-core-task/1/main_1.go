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

	fmt.Println(CheckType(intNum10))
	fmt.Println(CheckType(intNum8))
	fmt.Println(CheckType(intNum16))
	fmt.Println(CheckType(float64Num))
	fmt.Println(CheckType(stringSome))
	fmt.Println(CheckType(boolOp))
	fmt.Println(CheckType(complex64Num))

	combined := CombineVariablesToString(intNum10, intNum8, intNum16, float64Num, stringSome, boolOp, complex64Num)
	fmt.Println(combined)

	var runeOfSlice []rune = ConvertStringToSliceOfRune(combined)
	hashed := HashRunes(runeOfSlice)
	fmt.Println(hashed)
}

func CheckType(something any) string {
	return reflect.TypeOf(something).String()
}

func CombineVariablesToString(intNum10, intNum8, intNum16 int, float64Num float64, stringSome string, boolOp bool, complex64Num complex64) string {
	return fmt.Sprintf("%d %d %d %.2f %s %t %v", intNum10, intNum8, intNum16, float64Num, stringSome, boolOp, complex64Num)
}

func ConvertStringToSliceOfRune(text string) []rune {
	return []rune(text)
}

func HashRunes(something []rune) string {
	salt := []rune("go-2024")
	mid := len(something) / 2
	somethingSalt := append(something[:mid], append(salt, something[mid:]...)...)
	hash := sha256.Sum256([]byte(string(somethingSalt)))
	return hex.EncodeToString(hash[:])
}
