package leetcode

import (
	"fmt"
	"math"
)

func KnightProbability() {
	{
		n := 3
		k := 2
		row := 0
		column := 0
		fmt.Println(knightProbability(n, k, row, column)) // 0.06250
	}

	{
		n := 1
		k := 0
		row := 0
		column := 0
		fmt.Println(knightProbability(n, k, row, column)) // 1
	}
}

// bottom up
func knightProbability(n int, k int, r int, c int) float64 {
	DIRECTIONS := [8][2]int{
		{-1, -2},
		{-2, -1},
		{1, -2},
		{2, -1},
		{2, 1},
		{1, 2},
		{-1, 2},
		{-2, 1},
	}
	cache := make(map[[3]int]float64, int(math.Pow(8, math.Min(3, float64(k)))))
	cache[[3]int{0, r, c}] = 1
	for step := 1; step <= k; step++ {
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				for _, dir := range DIRECTIONS {
					prevrow := row + dir[0]
					prevcol := col + dir[1]
					if prevrow >= 0 && prevrow < n && prevcol >= 0 && prevcol < n {
						cache[[3]int{step, row, col}] += cache[[3]int{step - 1, prevrow, prevcol}] / 8
					}
				}
			}
		}
	}

	var res float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += cache[[3]int{k, i, j}]
		}
	}
	return res
}

// top down
// func knightProbability(n int, k int, row int, column int) float64 {
// 	DIRECTIONS := [8][2]int{
// 		{-1, -2},
// 		{-2, -1},
// 		{1, -2},
// 		{2, -1},
// 		{2, 1},
// 		{1, 2},
// 		{-1, 2},
// 		{-2, 1},
// 	}
// 	cache := make(map[[3]int]float64, int(math.Pow(8, math.Min(3, float64(k)))))
// 	var recurse func(k, r, c int) float64
// 	recurse = func(k, r, c int) float64 {
// 		if r < 0 || r >= n || c < 0 || c >= n {
// 			return 0
// 		}
// 		if k == 0 {
// 			return 1
// 		}
// 		if v, ok := cache[[3]int{k, r, c}]; ok {
// 			return v
// 		}
// 		var res float64
// 		for _, dir := range DIRECTIONS {
// 			res += recurse(k-1, r+dir[0], c+dir[1]) / 8
// 		}
// 		cache[[3]int{k, r, c}] = res
// 		return res
// 	}
// 	return recurse(k, row, column)
// }
