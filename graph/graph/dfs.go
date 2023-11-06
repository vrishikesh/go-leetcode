package graph

import (
	"container/list"
	"fmt"
)

func DfsOfGraph() {
	fmt.Println(dfsOfGraph(5, [][]int{
		{1, 2, 4},
		{0},
		{0},
		{4},
		{0, 3},
	}))
}

func dfsOfGraph(V int, adj [][]int) []int {
	v := make([]int, V)
	q := list.New()

	dfs(0, adj, v, q)
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	return []int{}
}

func dfs(node int, adj [][]int, v []int, q *list.List) {
	v[node] = 1
	q.PushBack(node)

	for _, vertex := range adj[node] {
		if v[vertex] == 0 {
			dfs(vertex, adj, v, q)
		}
	}
}
