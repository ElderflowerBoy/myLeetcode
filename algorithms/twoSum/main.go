package main

import "fmt"

func main() {
	var a = []int{2, 7, 11, 15}
	var target = 9

	fmt.Println(twoSum(a, target))
}

func twoSum(nums []int, target int) []int {
	var result []int
	for i, v := range nums {
		for j, k := range nums {
			if i == j {
				continue
			}
			if v+k == target {
				result = append(result, i, j)
				return result
			}
		}
	}

	return result
}
