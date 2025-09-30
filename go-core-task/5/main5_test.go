package main

import (
	"slices"
	"testing"
)

func TestHasIntersections(t *testing.T) {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}
	expected := []int{64, 3}
	ok, result := hasIntersections(slice1, slice2)
	if !ok {
		t.Errorf("expected true, but got false")
	}
	if !slices.Equal(expected, result) {
		t.Errorf("must be equal. expected %v, but got %v", expected, result)
	}
}

func TestHasIntersections_NotFound(t *testing.T) {
	slice1 := []int{65, 58, 678}
	slice2 := []int{64, 2, 3, 43}
	expected := []int{}
	ok, result := hasIntersections(slice1, slice2)
	if ok {
		t.Errorf("expected false, but got true")
	}
	if !slices.Equal(expected, result) {
		t.Errorf("must be equal. expected %v, but got %v", expected, result)
	}
}
