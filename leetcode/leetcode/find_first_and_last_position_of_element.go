package leetcode

import "fmt"

func SearchRange() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(array []int, target int) []int {
	n := len(array)
	left, right := 0, n-1
	x := binarySearch(array, target, left, right)
	if x == -1 {
		return []int{-1, -1}
	}

	tempL, tempR := x, x
	for {
		index := binarySearch(array, target, left, tempR-1)
		if index == -1 {
			break
		} else {
			tempR = index
		}
	}
	for {
		index := binarySearch(array, target, tempL+1, right)
		if index == -1 {
			break
		} else {
			tempL = index
		}
	}

	return []int{tempR, tempL}
}

func binarySearch(array []int, target, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if target < array[mid] {
			right = mid - 1
		} else if target > array[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
