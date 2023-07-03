package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func PermutationSequence(n, k int) string {
	var out strings.Builder
	numbers := []int{}
	fact := 1
	k = k - 1

	for i := 1; i < n; i++ {
		fact = fact * i
		numbers = append(numbers, i)
	}
	numbers = append(numbers, n)

	for {
		index := k / fact
		out.WriteString(strconv.Itoa(numbers[index]))
		numbers = append(numbers[0:index], numbers[index+1:]...)
		if len(numbers) == 0 {
			break
		}
		k = k % fact
		fact = fact / len(numbers)
	}

	fmt.Println(out.String())
	return out.String()
}
