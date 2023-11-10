package leetcode

import (
	"fmt"
	"math"
)

func IsValidBST() {
	{
		two := &TreeNode[int]{Val: 2}
		one := &TreeNode[int]{Val: 1}
		three := &TreeNode[int]{Val: 3}

		two.Left = one
		two.Right = three

		fmt.Println(isValidBST(two))
	}

	{
		one := &TreeNode[int]{Val: 1}
		three := &TreeNode[int]{Val: 3}
		four := &TreeNode[int]{Val: 4}
		five := &TreeNode[int]{Val: 5}
		six := &TreeNode[int]{Val: 6}

		five.Left = one
		five.Right = four
		four.Left = three
		four.Right = six

		fmt.Println(isValidBST(five))
	}

	{
		one := &TreeNode[int]{Val: 1}

		fmt.Println(isValidBST(one))
	}

	{
		fmt.Println(isValidBST(nil))
	}

	{
		two1 := &TreeNode[int]{Val: 2}
		two2 := &TreeNode[int]{Val: 2}
		two3 := &TreeNode[int]{Val: 2}

		two1.Left = two2
		two1.Right = two3

		fmt.Println(isValidBST(two1))
	}
}

func isValidBST(root *TreeNode[int]) bool {
	if root == nil {
		return true
	}

	prev := math.MinInt
	isValid := true
	var dfs func(node *TreeNode[int])
	dfs = func(node *TreeNode[int]) {
		if !isValid {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if prev >= node.Val {
			isValid = false
			return
		}
		prev = node.Val
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)

	return isValid
}
