package main

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"testing"
)

func TestCheckType(t *testing.T) {
	if res := CheckType(42); res != "int" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := CheckType(052); res != "int" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := CheckType(0x2A); res != "int" {
		t.Errorf("expected be int, but got %s", res)
	}
	var float32Num float64 = 53292.523
	if res := CheckType(float32Num); res != "float64" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := CheckType("hello"); res != "string" {
		t.Errorf("expected be int, but got %s", res)
	}
	if res := CheckType(true); res != "bool" {
		t.Errorf("expected be int, but got %s", res)
	}
	var cmp complex64 = complex(5, 2)
	if res := CheckType(cmp); res != "complex64" {
		t.Errorf("expected be int, but got %s", res)
	}
}

func TestCombineVariablesToString(t *testing.T) {
	result := CombineVariablesToString(42, 0o52, 0x2A, 3.14, "hello", true, 5+2i)
	expected := "42 42 42 3.14 hello true (5+2i)"
	if result != expected {
		t.Errorf("expected %q, but got %q", expected, result)
	}
}

func TestConvertStringToSliceOfRune(t *testing.T) {
	expectedRunes := []rune("hello")
	if result := ConvertStringToSliceOfRune("hello"); !slices.Equal(expectedRunes, result) {
		t.Errorf("expected %v, but got %v", expectedRunes, result)
	}
}

func TestHashRunes(t *testing.T) {
	textExemp := "my-text--end-of1"
	hash := sha256.Sum256([]byte("my-text-go-2024-end-of1"))
	expectedHashString := hex.EncodeToString(hash[:])

	if result := HashRunes([]rune(textExemp)); expectedHashString != result {
		t.Errorf("expected %s, but got %s", expectedHashString, result)
	}

}
