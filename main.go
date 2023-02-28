package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 2, 3}
	sum := Sum(numbers)
	max := Max(numbers)
	min := Min(numbers)
	avg := Avg(numbers)
	product := Product(numbers)
	count := Count(numbers)
	fmt.Println(sum, max, min, avg, product, count)
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
func Product(numbers []int) int {
	var result int = 1
	for _, number := range numbers {
		result = result * number
	}
	return result
}

func Count(numbers []int) int {
	var result int
	for _, _ = range numbers {
		result = result + 1
	}
	return result
}
