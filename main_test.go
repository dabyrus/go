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
	sum = Sum([]int{})
	if sum != 0 {
		t.Fatalf("Должно было быть %d, а было %d", 0, sum)
	}
}
func TestMax(t *testing.T) {
	max := Max([]int{1, 2, 3})
	if max != 3 {
		t.Fatalf("Должно было быть %d, а было %d", 3, max)
	}
	max = Max([]int{4, 2, 3})
	if max != 4 {
		t.Fatalf("Должно было быть %d, а было %d", 4, max)
	}
	max = Max([]int{-3, -2, -1})
	if max != -1 {
		t.Fatalf("Должно было быть %d, а было %d", -1, max)
	}
	max = Sum([]int{})
	if max != 0 {
		t.Fatalf("Должно было быть %d, а было %d", 0, max)
	}
}

func TestMin(t *testing.T) {
	max := Min([]int{1, 2, 3})
	if max != 1 {
		t.Fatalf("Должно было быть %d, а было %d", 1, max)
	}
	max = Min([]int{4, 2, 3})
	if max != 2 {
		t.Fatalf("Должно было быть %d, а было %d", 2, max)
	}
	max = Min([]int{-3, -2, -1})
	if max != -3 {
		t.Fatalf("Должно было быть %d, а было %d", -3, max)
	}
}
