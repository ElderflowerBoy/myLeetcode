package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(searchInsert(nums, 12))

}

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		avg := (min + max) / 2

		if nums[avg] < target {
			min = avg + 1
		} else {
			max = avg - 1
		}

		if nums[avg] == target {
			return avg
		}

	}

	return min
}
