package leetcode

import "fmt"

func LongestCommonSubsequence() {
	fmt.Println(longestCommonSubsequence("abcde", "ace")) // 3
	fmt.Println(longestCommonSubsequence("abc", "abc"))   // 3
	fmt.Println(longestCommonSubsequence("abc", "def"))   // 0
	fmt.Println(longestCommonSubsequence("abcde", "aec")) // 2
}

// bottom up
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	L := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		L[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				L[i][j] = 1 + L[i+1][j+1]
			} else {
				L[i][j] = max(L[i+1][j], L[i][j+1])
			}
		}
	}
	return L[0][0]
}

// top down
// func longestCommonSubsequence(text1 string, text2 string) int {
// 	m := len(text1)
// 	n := len(text2)
// 	temp := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		temp[i] = -1
// 	}

// 	L := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		L[i] = make([]int, n)
// 		copy(L[i], temp)
// 	}

// 	var lcsLength func(i, j int) int
// 	lcsLength = func(i, j int) int {
// 		if i == m || j == n {
// 			return 0
// 		}
// 		if L[i][j] < 0 {
// 			if text1[i] == text2[j] {
// 				L[i][j] = 1 + lcsLength(i+1, j+1)
// 			} else {
// 				L[i][j] = max(
// 					lcsLength(i+1, j),
// 					lcsLength(i, j+1),
// 				)
// 			}
// 		}
// 		return L[i][j]
// 	}
// 	return lcsLength(0, 0)
// }
