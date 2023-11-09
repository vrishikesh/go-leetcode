package leetcode

import "fmt"

func LevelOrder() {
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

		fmt.Println(levelOrder(three))
	}

	{
		one := &TreeNode[int]{Val: 1}

		fmt.Println(levelOrder(one))
	}

	{
		fmt.Println(levelOrder(nil))
	}
}

func levelOrder(root *TreeNode[int]) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	q := &Queue[*TreeNode[int]]{}
	q.Enqueue(root)

	for q.Len() > 0 {
		n := q.Len()
		curr_ans := []int{}
		for i := 0; i < n; i++ {
			node := q.Dequeue()
			curr_ans = append(curr_ans, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		ans = append(ans, curr_ans)
	}

	return ans
}
