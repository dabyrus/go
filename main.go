package main

import (
	"fmt"
)

func main() {
	numbers := []int{3, 4, 9, 6, 9, 99, 1, 8, 45}
	sum := Sum(numbers)
	max := Max(numbers)
	min := Min(numbers)
	avg := Avg(numbers)
	fmt.Println(sum, max, min, avg)
}

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result = result + number
	}
	return result
}
func Max(numbers []int) int {
	var result int = numbers[0]
	for _, n := range numbers {
		if result < n {
			result = n
		}
	}
	return result
}

func Min(numbers []int) int {
	var result int = numbers[0]
	for _, n := range numbers {
		if result > n {
			result = n
		}
	}
	return result
}

func Avg(numbers []int) int {
	return Sum(numbers) / len(numbers)
}
//нужно перемножить
func Product(numbers []int) int {
	return 0
}

func Count(numbers []int) int {
	var result int
	for _, _ = range numbers {
		result = result + 1
	}
	return result
}
