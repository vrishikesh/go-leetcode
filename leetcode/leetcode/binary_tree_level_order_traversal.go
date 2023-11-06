package leetcode

import "fmt"

type LO_TreeNode struct {
	Val   int
	Left  *LO_TreeNode
	Right *LO_TreeNode
}

type LO_Queue []*LO_TreeNode

func (q *LO_Queue) Len() int {
	return len(*q)
}

func (q *LO_Queue) Enqueue(x *LO_TreeNode) {
	*q = append(*q, x)
}

func (q *LO_Queue) Dequeue() *LO_TreeNode {
	cp := *q
	x := cp[0]
	*q = cp[1:]
	return x
}

func LevelOrder() {
	{
		three := &LO_TreeNode{Val: 3}
		nine := &LO_TreeNode{Val: 9}
		twenty := &LO_TreeNode{Val: 20}
		fifteen := &LO_TreeNode{Val: 15}
		seven := &LO_TreeNode{Val: 7}

		three.Left = nine
		three.Right = twenty
		twenty.Left = fifteen
		twenty.Right = seven

		fmt.Println(levelOrder(three))
	}

	{
		one := &LO_TreeNode{Val: 1}

		fmt.Println(levelOrder(one))
	}

	{
		fmt.Println(levelOrder(nil))
	}
}

func levelOrder(root *LO_TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	q := &LO_Queue{}
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
