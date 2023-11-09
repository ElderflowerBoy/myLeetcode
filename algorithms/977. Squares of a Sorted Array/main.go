package main

func main() {

}
func sortedSquares(nums []int) []int {
	l := len(nums) - 1
	res := make([]int, len(nums))

	for i := range nums {
		nums[i] *= nums[i]
	}

	head := 0
	tail := l

	for l >= 0 {
		if nums[head] > nums[tail] {
			res[l] = nums[head]
			head++
		} else {
			res[l] = nums[tail]
			tail--
		}
		l--
	}

	return res

}
