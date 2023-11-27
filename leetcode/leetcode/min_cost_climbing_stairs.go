package leetcode

import (
	"fmt"
)

func MinCostClimbingStairs() {
	{
		cost := []int{10, 15, 20}
		fmt.Println(minCostClimbingStairs(cost))
	}

	{
		cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
		fmt.Println(minCostClimbingStairs(cost))
	}
}

// bottom up
func minCostClimbingStairs(cost []int) int {
	N := len(cost)
	first, second := 0, cost[0]
	for i := 1; i < N; i++ {
		first, second = second, cost[i]+min(first, second)
	}
	return min(first, second)
}

// top down
// func minCostClimbingStairs(cost []int) int {
// 	N := len(cost)
// 	var minCost func(idx int) int
// 	cache := make(map[int]int, N)
// 	minCost = func(idx int) int {
// 		if idx < 0 {
// 			return 0
// 		}
// 		if idx <= 1 {
// 			return cost[idx]
// 		}
// 		if _, ok := cache[idx]; !ok {
// 			cache[idx] = cost[idx] + min(minCost(idx-1), minCost(idx-2))
// 		}
// 		return cache[idx]
// 	}
// 	return min(minCost(N-1), minCost(N-2))
// }

// naive recursion
// func minCostClimbingStairs(cost []int) int {
// 	minCost := math.MaxInt
// 	N := len(cost)
// 	var dfs func(currentCost, idx int)
// 	dfs = func(currentCost, idx int) {
// 		if idx+1 >= N || idx+2 >= N {
// 			minCost = min(minCost, currentCost)
// 			return
// 		}
// 		dfs(currentCost+cost[idx+1], idx+1)
// 		dfs(currentCost+cost[idx+2], idx+2)
// 	}
// 	dfs(0, -1)
// 	return minCost
// }
