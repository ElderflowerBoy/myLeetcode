package main

import "fmt"

func main() {

	var a = []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(search(a, 12))
}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	min := 0
	max := len(nums) - 1

	for min <= max {
		median := (min + max) / 2

		if nums[median] < target {
			min = median + 1
		} else {
			max = median - 1
		}

		if nums[median] == target {
			return median
		}

	}

	return -1
}
