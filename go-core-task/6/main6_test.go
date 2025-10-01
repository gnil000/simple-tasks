package main

import "testing"

func TestRandNumsGenerator(t *testing.T) {
	expectedCount := 10
	result := randNumsGenerator(expectedCount)
	count := 0
	for range result {
		count++
	}
	if expectedCount != count {
		t.Errorf("expected count: %d, but got: %d", expectedCount, count)
	}
}
