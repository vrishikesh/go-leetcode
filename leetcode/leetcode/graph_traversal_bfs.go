package leetcode

import "fmt"

func GraphTraversalBFS() {
	{
		adj := [][]int{
			{1, 3},
			{0},
			{3, 8},
			{0, 2, 4, 5},
			{3, 6},
			{3},
			{4, 7},
			{6},
			{2},
		}
		fmt.Println(graphTraversalBFS(adj))
	}
}

func graphTraversalBFS(adj [][]int) []int {
	L := len(adj)
	values := make([]int, 0, L)
	visited := make([]bool, L)
	q := &Queue[int]{}
	q.Enqueue(0)

	for q.Len() > 0 {
		vertex := q.Dequeue()
		visited[vertex] = true
		list := adj[vertex]
		values = append(values, vertex)

		for _, v := range list {
			if !visited[v] {
				q.Enqueue(v)
			}
		}
	}

	return values
}
