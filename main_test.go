package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum([]int{1, 2, 3})
	if sum != 6 {
		t.Fatalf("Должно было быть %d, а было %d", 6, sum)
	}
	sum = Sum([]int{4, 2, 3})
	if sum != 9 {
		t.Fatalf("Должно было быть %d, а было %d", 9, sum)
	}
}
