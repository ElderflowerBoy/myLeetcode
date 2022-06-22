package main

func main() {
	println(removeElement([]int{3, 2, 2, 3}, 3))
	println(removeElement([]int{-11, -10, -10, 3, 3, 3, 3, 0, 0, 1, 2, 3, 3, 3}, 3))
	println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
