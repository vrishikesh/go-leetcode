package leetcode

import (
	"fmt"
)

func FindKthLargest() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	indexToFind := n - k
	low, high := 0, n-1

	for {
		pi := partition(nums, low, high)
		if pi == indexToFind {
			return nums[pi]
		} else if pi < indexToFind {
			low = pi + 1
		} else {
			high = pi - 1
		}
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]

	i := low
	for j := low; j <= high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}
