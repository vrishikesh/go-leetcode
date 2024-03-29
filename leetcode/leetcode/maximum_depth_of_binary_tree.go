package leetcode

import "fmt"

func MaxDepth() {
	{
		three := &TreeNode[int]{Val: 3}
		nine := &TreeNode[int]{Val: 9}
		twenty := &TreeNode[int]{Val: 20}
		fifteen := &TreeNode[int]{Val: 15}
		seven := &TreeNode[int]{Val: 7}

		three.Left = nine
		three.Right = twenty
		twenty.Left = fifteen
		twenty.Right = seven

		fmt.Println(maxDepth(three))
	}

	{
		one := &TreeNode[int]{Val: 1}
		two := &TreeNode[int]{Val: 2}

		one.Right = two

		fmt.Println(maxDepth(one))
	}

	{
		fmt.Println(maxDepth(nil))
	}

	{
		one := &TreeNode[int]{Val: 1}
		fmt.Println(maxDepth(one))
	}
}

func maxDepth(root *TreeNode[int]) int {
	if root == nil {
		return 0
	}
	maximumDepth := 1

	var dfs func(*TreeNode[int], int)
	dfs = func(node *TreeNode[int], depth int) {
		if node == nil {
			return
		}
		if depth > maximumDepth {
			maximumDepth = depth
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)

	return maximumDepth
}
