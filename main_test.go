package main

import "testing"

func TestSum(t *testing.T) {
	var numbers []int = []int{1, 2, 3, 4}

	var expected int = 10
	var actual int = sum(numbers)

	if expected != actual {
		t.Errorf("expected: %v got: %v\n", expected, actual)
	}
}
