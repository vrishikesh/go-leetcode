package leetcode

import "fmt"

func GetRow() {
	fmt.Println(getRow(3)) // 1 3 3 1
	fmt.Println(getRow(0)) // 1
	fmt.Println(getRow(1)) // 1 1
}

// top down
func getRow(rowIndex int) []int {
	X := make([]int, rowIndex+1)
	Y := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		X[0], X[i] = 1, 1
		for j := 1; j < i; j++ {
			X[j] = Y[j-1] + Y[j]
		}
		copy(Y, X)
	}
	return X
}

// top down
// func getRow(rowIndex int) []int {
// 	M := make([][]int, rowIndex+1)
// 	for i := 0; i < rowIndex+1; i++ {
// 		M[i] = make([]int, rowIndex+1)
// 	}
// 	var dfs func(row int) []int
// 	dfs = func(row int) []int {
// 		for j := 0; j <= row; j++ {
// 			if j == 0 || j == row {
// 				M[row][j] = 1
// 			} else if M[row-1][j-1] > 0 && M[row-1][j] > 0 {
// 				M[row][j] = M[row-1][j-1] + M[row-1][j]
// 			} else {
// 				A := dfs(row - 1)
// 				M[row][j] = A[j-1] + A[j]
// 			}
// 		}
// 		return M[row]
// 	}
// 	ans := dfs(rowIndex)
// 	fmt.Println(M)
// 	return ans
// }
