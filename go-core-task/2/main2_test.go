package main

import (
	"slices"
	"testing"
)

func TestCreateRandomSlice(t *testing.T) {
	slice1 := createRandomSlice()
	slice2 := createRandomSlice()
	if slices.Equal(slice1, slice2) {
		t.Error("expected different slices, but got similar")
	}
}

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedEvenSlice := []int{2, 4, 6, 8, 10}
	resultSlice := sliceExample(slice)
	if len(resultSlice) > len(slice) {
		t.Errorf("length new slice (len = %d) cannot be more than original slice (len = %d)", len(resultSlice), len(slice))
	}
	if !slices.Equal(expectedEvenSlice, resultSlice) {
		t.Errorf("expected %v, but got %v", expectedEvenSlice, resultSlice)
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 112}
	numForAdd := 112
	result := addElements(slice, numForAdd)
	if !slices.Equal(expectedSlice, result) {
		t.Errorf("expected %v, but got %v", expectedSlice, result)
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := copySlice(slice)
	if !slices.Equal(expectedSlice, result) {
		t.Errorf("must be equal. expected %v, but got %v", expectedSlice, result)
	}
	result[0] = 666
	if slices.Equal(expectedSlice, result) {
		t.Errorf("shouldn't be equal. expected %v, but got %v", expectedSlice, result)
	}
}

func TestRemoveElement(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected1 := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected2 := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	result1 := removeElement(slice1, 0)
	if !slices.Equal(expected1, result1) {
		t.Errorf("must be equal. expected %v, but got %v", expected1, result1)
	}
	result2 := removeElement(slice2, 5)
	if !slices.Equal(expected2, result2) {
		t.Errorf("must be equal. expected %v, but got %v", expected2, result2)
	}
}
