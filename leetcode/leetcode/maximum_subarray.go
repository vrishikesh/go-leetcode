package leetcode

import "fmt"

func MaxSubArray() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

// bottom up
func maxSubArray(nums []int) int {
	N := len(nums)
	currentSum, largest := nums[0], nums[0]
	for i := 1; i < N; i++ {
		currentSum = max(nums[i]+currentSum, nums[i])
		largest = max(largest, currentSum)
	}
	return largest
}

// recurrence relation - top down
// func maxSubArray(nums []int) int {
// 	N := len(nums)
// 	cache := map[[2]int]int{}
// 	var recurse func(i, sum int) int
// 	recurse = func(i, sum int) int {
// 		if i >= N {
// 			return sum
// 		}
// 		if _, ok := cache[[2]int{i, sum}]; !ok {
// 			cache[[2]int{i, sum}] = max(
// 				recurse(i+1, nums[i]+sum),
// 				recurse(i+1, nums[i]),
// 			)
// 		}
// 		return cache[[2]int{i, sum}]
// 	}
// 	return recurse(1, nums[0])
// }

// naive
// func maxSubArray(nums []int) int {
// 	N := len(nums)
// 	largest := 0
// 	for i := 0; i < N; i++ {
// 		currentSum := 0
// 		for j := i; j < N; j++ {
// 			currentSum += nums[j]
// 			largest = max(largest, currentSum)
// 		}
// 	}
// 	return largest
// }
