package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
	var n = []int{-1, -100, 3, 99}
	rotate(n, 2)
	fmt.Println(n)

}

func rotate(nums []int, k int) {
	var length = len(nums)
	var rotates = length - k%length
	var res = make([]int, length)
	if length > 1 && rotates > 0 {
		res = append(nums[rotates:], nums[:rotates]...)
		for i := 0; i < length; i++ {
			nums[i] = res[i]
		}
	}
}
