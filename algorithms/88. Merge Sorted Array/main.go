package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 2, []int{1, 5, 6}, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var bufArray, j, k = make([]int, m+n, m+n), 0, 0
	copy(bufArray, nums1)
	if len(nums2) < 1 {
		return
	}
	for i := 0; i < m+n; i++ {
		if j < m && (k >= n || bufArray[j] <= nums2[k]) {
			nums1[i] = bufArray[j]
			j++
			continue
		}

		if k < n {
			nums1[i] = nums2[k]
			k++
		}
	}
}
