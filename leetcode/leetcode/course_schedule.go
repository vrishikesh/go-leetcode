package leetcode

import "fmt"

func CanFinish() {
	{
		numCourses := 2
		prerequisites := [][]int{
			{1, 0},
		}
		fmt.Println(canFinish(numCourses, prerequisites))
	}

	{
		numCourses := 2
		prerequisites := [][]int{
			{1, 0},
			{0, 1},
		}
		fmt.Println(canFinish(numCourses, prerequisites))
	}

	{
		numCourses := 6
		prerequisites := [][]int{
			{1, 0},
			{2, 1},
			{2, 5},
			{0, 3},
			{4, 3},
			{3, 5},
			{4, 5},
		}
		fmt.Println(canFinish(numCourses, prerequisites))
	}

	{
		numCourses := 7
		prerequisites := [][]int{
			{0, 3},
			{1, 0},
			{2, 1},
			{4, 5},
			{6, 4},
			{5, 6},
		}
		fmt.Println(canFinish(numCourses, prerequisites))
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		indegree[p[0]] += 1
		adj[p[1]] = append(adj[p[1]], p[0])
	}

	q := &Queue[int]{}
	for v, d := range indegree {
		if d == 0 {
			q.Enqueue(v)
		}
	}

	count := 0
	for q.Len() > 0 {
		v := q.Dequeue()
		count += 1
		for _, a := range adj[v] {
			indegree[a] -= 1
			if indegree[a] == 0 {
				q.Enqueue(a)
			}
		}
	}

	return count == numCourses
}

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	adj := make([][]int, numCourses)
// 	for _, p := range prerequisites {
// 		adj[p[1]] = append(adj[p[1]], p[0])
// 	}
// 	fmt.Println(adj)

// 	for v := 0; v < numCourses; v++ {
// 		q := &Queue[int]{}
// 		visited := make([]bool, numCourses)
// 		q.Enqueue(v)
// 		for q.Len() > 0 {
// 			vertex := q.Dequeue()
// 			visited[vertex] = true
// 			for _, conn := range adj[vertex] {
// 				if v == conn {
// 					return false
// 				}
// 				if !visited[conn] {
// 					q.Enqueue(conn)
// 				}
// 			}
// 		}
// 	}

// 	return true
// }
