package leetcode

import (
	"fmt"
	"math"
)

func CountNodes() {
	{
		one := &TreeNode{Val: 1}
		two := &TreeNode{Val: 2}
		three := &TreeNode{Val: 3}
		four := &TreeNode{Val: 4}
		five := &TreeNode{Val: 5}
		six := &TreeNode{Val: 6}

		one.Left = two
		one.Right = three
		two.Left = four
		two.Right = five
		three.Left = six

		fmt.Println(countNodes(one)) // 6
	}

	{
		fmt.Println(countNodes(nil)) // 0
	}

	{
		one := &TreeNode{Val: 1}

		fmt.Println(countNodes(one)) // 1
	}

	{
		one := &TreeNode{Val: 1}
		two := &TreeNode{Val: 2}
		three := &TreeNode{Val: 3}
		four := &TreeNode{Val: 4}

		one.Left = two
		one.Right = three
		two.Left = four

		fmt.Println(countNodes(one)) // 4
	}
}

/** Striver's Solution */
// func countNodes(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	l := getLeftHeight(node)
// 	r := getRightHeight(node)
// 	if l == r {
// 		return int(math.Pow(2, float64(l+1))) - 1
// 	}
// 	return 1 + countNodes(node.Left) + countNodes(node.Right)
// }

// func getLeftHeight(node *TreeNode) int {
// 	height := 0
// 	for node.Left != nil {
// 		height += 1
// 		node = node.Left
// 	}
// 	return height
// }

// func getRightHeight(node *TreeNode) int {
// 	height := 0
// 	for node.Right != nil {
// 		height += 1
// 		node = node.Right
// 	}
// 	return height
// }

/** Solution from Udemy */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := getTreeHeight(root)
	upperCount := int(math.Pow(2, float64(height))) - 1
	left, right := 0, upperCount
	for left < right {
		idxToFind := int(math.Ceil(float64(left+right) / 2))
		if nodeExists(idxToFind, height, root) {
			left = idxToFind
		} else {
			right = idxToFind - 1
		}
	}

	return upperCount + left + 1
}

func nodeExists(idxToFind int, height int, node *TreeNode) bool {
	left, right, count := 0, int(math.Pow(2, float64(height)))-1, 0
	for count < height {
		mid := int(math.Ceil(float64(left+right) / 2))
		if idxToFind >= mid {
			node = node.Right
			left = mid
		} else {
			node = node.Left
			right = mid - 1
		}
		count += 1
	}
	return node != nil
}

func getTreeHeight(node *TreeNode) int {
	height := 0
	for node.Left != nil {
		height += 1
		node = node.Left
	}
	return height
}

/** My Initial Solution */
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	curr := root
// 	height := 0

// 	for curr != nil {
// 		height += 1
// 		curr = curr.Left
// 	}

// 	noOfNodes := int(math.Pow(2, float64(height))) - 1
// 	depths := 0
// 	found := false
// 	var dfs func(node *TreeNode, depth, position int)
// 	dfs = func(node *TreeNode, depth, position int) {
// 		if found {
// 			return
// 		}
// 		if node == nil {
// 			noOfNodes -= 1
// 			return
// 		}
// 		if depths == depth {
// 			depths += 1
// 		}
// 		if height == depth+1 {
// 			found = true
// 			return
// 		}
// 		dfs(node.Right, depth+1, position)
// 		dfs(node.Left, depth+1, position)
// 	}
// 	dfs(root, 0, 0)

// 	return noOfNodes
// }
