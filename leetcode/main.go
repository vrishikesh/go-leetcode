package main

import "github.com/vrishikesh/go-leetcode/leetcode/leetcode"

func main() {
	// leetcode.PrintAllSubsequences([]int{1, 2, 1})
	// leetcode.PrintAllSSThatSumToK([]int{1, 2, 1}, 2)
	// leetcode.PrintTheFirstSSThatSumToK([]int{1, 2, 1}, 2)
	// leetcode.PrintTheCountOfSsThatSumToK([]int{1, 2, 1}, 2)
	// leetcode.PrintAllCombSubsequencesThatSumToK([]int{2, 3, 6, 7}, 7)
	// leetcode.NumSquaresDFS(12)
	// leetcode.NumSquaresDP(12)
	// leetcode.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	// leetcode.PrintAllPermutationsMap([]int{1, 2, 3})
	// leetcode.PrintAllPermutationsSwap([3]int{1, 2, 3})
	// leetcode.NQueens(4)
	// leetcode.SolveSudoku([][]byte{
	// 	{5, 3, '.', '.', 7, '.', '.', '.', '.'},
	// 	{6, '.', '.', 1, 9, 5, '.', '.', '.'},
	// 	{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
	// 	{8, '.', '.', '.', 6, '.', '.', '.', 3},
	// 	{4, '.', '.', 8, '.', 3, '.', '.', 1},
	// 	{7, '.', '.', '.', 2, '.', '.', '.', 6},
	// 	{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
	// 	{'.', '.', '.', 4, 1, 9, '.', '.', 5},
	// 	{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	// })
	// leetcode.MColoring(4, 3, 5, [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}})
	// leetcode.PalindromePartitioning("aab")
	leetcode.RatInAMaze([][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	})
}
