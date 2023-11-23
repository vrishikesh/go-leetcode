package leetcode

import "fmt"

func GraphTraversalDFS() {
	adj := [][]int{
		{1, 3},
		{0},
		{3, 8},
		{0, 4, 5, 2},
		{3, 6},
		{3},
		{4, 7},
		{6},
		{2},
	}
	fmt.Println(graphTraversalDFS(adj))
}

func graphTraversalDFS(adj [][]int) []int {
	values := make([]int, 0, len(adj))
	visited := make([]bool, len(adj))

	var travelDFS func(vertex int)
	travelDFS = func(vertex int) {
		if visited[vertex] {
			return
		}

		values = append(values, vertex)
		visited[vertex] = true

		for _, v := range adj[vertex] {
			travelDFS(v)
		}
	}
	travelDFS(0)

	return values
}
