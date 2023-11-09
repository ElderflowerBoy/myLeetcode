package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 1}))
}

func containsDuplicate(nums []int) bool {
	var m = map[int]struct{}{}

	for _, i := range nums {
		if _, ok := m[i]; ok {
			return false
		} else {
			m[i] = struct{}{}
		}
	}
	return true
}
