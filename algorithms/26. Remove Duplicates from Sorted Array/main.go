package main

import "fmt"

func main() {
	println(removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4, 4, 5}))
	println(removeDuplicates([]int{1, 1, 2}))
	println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	println(removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 7}))
	println(removeDuplicates([]int{1}))
	println(removeDuplicates([]int{-11, -10, -10, 0, 0, 1, 2, 3, 3, 3, 3, 3, 3, 3}))

}

func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	fmt.Println(nums)
	return k + 1
}
