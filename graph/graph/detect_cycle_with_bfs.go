package graph

import (
	"fmt"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

type Node struct {
	Val, ParentVal int
}

func DetectCycleBFS(adj [][]int) {
	N := len(adj)
	v := make([]bool, N)
	isCycle := false

	for i := 0; i < N; i++ {
		if !v[i] {
			if detectCycle(i, adj, v) {
				isCycle = true
				break
			}
		}
	}

	fmt.Println(isCycle)
	fmt.Println(v)
}

func detectCycle(src int, adj [][]int, v []bool) bool {
	q := &queue.Queue[Node]{}
	q.Push(Node{Val: src, ParentVal: -1})
	v[src] = true

	for !q.IsEmpty() {
		node := q.Pop()
		list := adj[node.Val]
		// fmt.Printf("%+2v %+v\n", node, list)

		for _, l := range list {
			if !v[l] {
				q.Push(Node{Val: l, ParentVal: node.Val})
				v[l] = true
			} else if l != node.ParentVal {
				return true
			}
		}
	}

	return false
}
