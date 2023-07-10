package graph

import (
	"fmt"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

func DfsOfGraph(V int, adj [][]int) []int {
	v := make([]int, V)
	q := queue.Queue[int]{}

	dfs(0, adj, v, &q)
	fmt.Println(q.String())

	return []int{}
}

func dfs(node int, adj [][]int, v []int, q *queue.Queue[int]) {
	v[node] = 1
	q.Push(node)

	for _, vertex := range adj[node] {
		if v[vertex] == 0 {
			dfs(vertex, adj, v, q)
		}
	}
}
