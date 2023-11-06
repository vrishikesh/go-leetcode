package leetcode

import "fmt"

func NumberOfInversions() {
	numberOfInversions([]int{5, 4, 3, 2, 1})
}

func numberOfInversions(a []int) {
	N := len(a)
	temp := make([]int, N)
	copy(temp, a)

	fmt.Println(mergeSort(a, temp, 0, N-1))
	fmt.Println(a)
}

func mergeSort(a, temp []int, low, high int) int {
	count := 0
	if low >= high {
		return count
	}
	mid := (low + high) / 2

	count += mergeSort(a, temp, low, mid)
	count += mergeSort(a, temp, mid+1, high)
	count += merge(a, temp, low, mid, high)

	return count
}

func merge(a, temp []int, low, mid, high int) int {
	left := low
	right := mid + 1
	index := low
	count := 0

	for left <= mid && right <= high {
		if a[left] <= a[right] {
			temp[index] = a[left]
			left += 1
			index += 1
		} else {
			count += mid - left + 1
			temp[index] = a[right]
			right += 1
			index += 1
		}
	}

	for left <= mid {
		temp[index] = a[left]
		left += 1
		index += 1
	}

	for right <= high {
		temp[index] = a[right]
		right += 1
		index += 1
	}

	for i := low; i <= high; i++ {
		a[i] = temp[i]
	}

	return count
}
