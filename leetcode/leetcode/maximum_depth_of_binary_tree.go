package leetcode

import "fmt"

type MD_TreeNode struct {
	Val   int
	Left  *MD_TreeNode
	Right *MD_TreeNode
}

func MaxDepth() {
	{
		three := &MD_TreeNode{Val: 3}
		nine := &MD_TreeNode{Val: 9}
		twenty := &MD_TreeNode{Val: 20}
		fifteen := &MD_TreeNode{Val: 15}
		seven := &MD_TreeNode{Val: 7}

		three.Left = nine
		three.Right = twenty
		twenty.Left = fifteen
		twenty.Right = seven

		fmt.Println(maxDepth(three))
	}

	{
		one := &MD_TreeNode{Val: 1}
		two := &MD_TreeNode{Val: 2}

		one.Right = two

		fmt.Println(maxDepth(one))
	}

	{
		fmt.Println(maxDepth(nil))
	}

	{
		one := &MD_TreeNode{Val: 1}
		fmt.Println(maxDepth(one))
	}
}

func maxDepth(root *MD_TreeNode) int {
	if root == nil {
		return 0
	}
	maximumDepth := 1
	dfs(root, 1, &maximumDepth)
	return maximumDepth
}

func dfs(node *MD_TreeNode, depth int, maximumDepth *int) {
	if node == nil {
		return
	}
	if depth > *maximumDepth {
		*maximumDepth = depth
	}
	dfs(node.Left, depth+1, maximumDepth)
	dfs(node.Right, depth+1, maximumDepth)
}
