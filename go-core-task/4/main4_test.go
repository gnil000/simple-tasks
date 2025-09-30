package main

import (
	"slices"
	"testing"
)

func TestLeftJoinSlices_Ok(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	result := LeftJoinSlices(slice1, slice2)
	if !slices.Equal(expected, result) {
		t.Errorf("must be equal. expected %v, but got %v", expected, result)
	}
}

func TestLeftJoinSlices_EmptyBothSlices(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{}
	expected := []string{}
	result := LeftJoinSlices(slice1, slice2)
	if !slices.Equal(expected, result) {
		t.Errorf("must be equal. expected %v, but got %v", expected, result)
	}
}

func TestLeftJoinSlices_EmptyFirstSlice(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{}
	result := LeftJoinSlices(slice1, slice2)
	if !slices.Equal(expected, result) {
		t.Errorf("must be equal. expected %v, but got %v", expected, result)
	}
}

func TestLeftJoinSlices_EmptySecondSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{}
	expected := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	result := LeftJoinSlices(slice1, slice2)
	if !slices.Equal(expected, result) {
		t.Errorf("must be equal. expected %v, but got %v", expected, result)
	}
}
