package graph

import (
	"container/list"
	"fmt"
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
	q := list.New()
	q.PushBack(Node{Val: src, ParentVal: -1})
	v[src] = true

	for q.Len() > 0 {
		node := q.Front().Value.(Node)
		q.Remove(q.Front())
		edges := adj[node.Val]
		// fmt.Printf("%+2v %+v\n", node, list)

		for _, l := range edges {
			if !v[l] {
				q.PushBack(Node{Val: l, ParentVal: node.Val})
				v[l] = true
			} else if l != node.ParentVal {
				return true
			}
		}
	}

	return false
}
