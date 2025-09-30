package main

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"testing"
)

func TestCheckType(t *testing.T) {
	if res := checkType(42); res != "int" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := checkType(052); res != "int" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := checkType(0x2A); res != "int" {
		t.Errorf("expected be int, but got %s", res)
	}
	var float32Num float64 = 53292.523
	if res := checkType(float32Num); res != "float64" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := checkType("hello"); res != "string" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := checkType(true); res != "bool" {
		t.Errorf("expected be int, but got %s", res)
	}
	var cmp complex64 = complex(5, 2)
	if res := checkType(cmp); res != "complex64" {
		t.Errorf("expected be int, but got %s", res)
	}
}

func TestCombineVariablesToString(t *testing.T) {
	result := combineVariablesToString(42, 0o52, 0x2A, 3.14, "hello", true, 5+2i)
	expected := "42 42 42 3.14 hello true (5+2i)"
	if result != expected {
		t.Errorf("expected %q, but got %q", expected, result)
	}
}

func TestConvertStringToSliceOfRune(t *testing.T) {
	expectedRunes := []rune("hello")
	if result := convertStringToSliceOfRune("hello"); !slices.Equal(expectedRunes, result) {
		t.Errorf("expected %v, but got %v", expectedRunes, result)
	}
}

func TestHashRunes(t *testing.T) {
	textExemp := "my-text--end-of1"
	hash := sha256.Sum256([]byte("my-text-go-2024-end-of1"))
	expectedHashString := hex.EncodeToString(hash[:])

	if result := hashRunes([]rune(textExemp)); expectedHashString != result {
		t.Errorf("expected %s, but got %s", expectedHashString, result)
	}

}
