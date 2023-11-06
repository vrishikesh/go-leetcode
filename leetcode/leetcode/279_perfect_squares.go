package leetcode

import (
	"container/list"
	"fmt"
	"math"
)

func NumSquaresDFS() {
	fmt.Println(numSquaresDFS(12))
}

// DFS O(2^n)
func numSquaresDFS(n int) int {
	var f func(int, int, int) int
	q := list.New()

	f = func(val int, s int, times int) int {
		if val == 0 {
			if s == n {
				for e := q.Front(); e != nil; e = e.Next() {
					fmt.Printf("%v ", e.Value)
				}
				fmt.Println()
				if q.Len() < times {
					times = q.Len()
				}
			}
			return times
		}

		q.PushBack(val * val)
		if s+val*val <= n {
			if t := f(val, s+val*val, times); t < times {
				times = t
			}
		}

		q.Remove(q.Front())
		if t := f(val-1, s, times); t < times {
			times = t
		}

		return times
	}

	max := sqrt(n)
	return f(max, 0, 10000)
}

func sqrt(n int) int {
	y := float64(n)
	x := math.Sqrt(y)
	return int(x)
}

func NumSquaresDP() {
	fmt.Println(numSquaresDP(12))
}

// DP O(n * n^1/2)
func numSquaresDP(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 1; i < n+1; i++ {
		minSquares := i
		for j := 1; j < i+1; j++ {
			square := j * j
			if i-square < 0 {
				break
			}
			minSquares = min(minSquares, 1+dp[i-square])
		}
		dp[i] = minSquares
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
