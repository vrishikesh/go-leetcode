package leetcode

import "fmt"

func RightSideView() {
	{
		one := &TreeNode{Val: 1}
		two := &TreeNode{Val: 2}
		three := &TreeNode{Val: 3}
		four := &TreeNode{Val: 4}
		five := &TreeNode{Val: 5}
		six := &TreeNode{Val: 6}
		seven := &TreeNode{Val: 7}
		eight := &TreeNode{Val: 8}

		one.Left = two
		one.Right = three
		two.Left = four
		two.Right = five
		three.Right = six
		four.Right = seven
		seven.Left = eight

		fmt.Println(rightSideView(one)) // [1, 3, 6, 7, 8]
		// Pre:  1 3 6 2 5 4 7 8
		// In:   6 3 1 5 2 7 8 4
		// Post: 6 3 5 8 7 4 2 1
	}

	{
		one := &TreeNode{Val: 1}
		two := &TreeNode{Val: 2}
		three := &TreeNode{Val: 3}
		five := &TreeNode{Val: 5}
		four := &TreeNode{Val: 4}

		one.Left = two
		one.Right = three
		two.Right = five
		three.Right = four

		fmt.Println(rightSideView(one)) // [1, 3, 4]
	}

	{
		one := &TreeNode{Val: 1}
		three := &TreeNode{Val: 3}

		one.Right = three

		fmt.Println(rightSideView(one)) // [1, 3]
	}

	{
		one := &TreeNode{Val: 1}
		two := &TreeNode{Val: 2}

		one.Left = two

		fmt.Println(rightSideView(one)) // [1, 2]
	}

	{
		one := &TreeNode{Val: 1}

		fmt.Println(rightSideView(one)) // [1]
	}

	{
		fmt.Println(rightSideView(nil)) // []
	}
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	// var bfs = func(root *TreeNode) {
	// 	q := &Queue[*TreeNode]{}
	// 	q.Enqueue(root)
	// 	for q.Len() > 0 {
	// 		n := q.Len()
	// 		for i := 0; i < n; i++ {
	// 			node := q.Dequeue()
	// 			if i+1 == n {
	// 				ans = append(ans, node.Val)
	// 			}
	// 			if node.Left != nil {
	// 				q.Enqueue(node.Left)
	// 			}
	// 			if node.Right != nil {
	// 				q.Enqueue(node.Right)
	// 			}
	// 		}
	// 	}
	// }

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(ans) == depth {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}

	// bfs(root)
	dfs(root, 0)
	return ans
}
