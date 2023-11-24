package leetcode

import (
	"fmt"
	"math"
)

func NetworkDelayTime() {
	{
		times := [][]int{
			{2, 1, 1},
			{2, 3, 1},
			{3, 4, 1},
		}
		n := 4
		k := 2
		fmt.Println(networkDelayTime(times, n, k)) // 2
	}

	{
		times := [][]int{
			{1, 2, 1},
		}
		n := 2
		k := 1
		fmt.Println(networkDelayTime(times, n, k)) // 1
	}

	{
		times := [][]int{
			{1, 2, 1},
		}
		n := 2
		k := 2
		fmt.Println(networkDelayTime(times, n, k)) // -1
	}

	{
		times := [][]int{
			{1, 2, 9},
			{1, 4, 2},
			{2, 5, 1},
			{4, 2, 4},
			{4, 5, 6},
			{3, 2, 3},
			{5, 3, 7},
			{3, 1, 5},
		}
		n := 5
		k := 1
		fmt.Println(networkDelayTime(times, n, k)) // 14
	}
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][][2]int, n+1)
	for _, v := range times {
		adj[v[0]] = append(adj[v[0]], [2]int{v[1], v[2]})
	}
	distances := make([]int, n+1)
	for i := 1; i <= n; i++ {
		distances[i] = math.MaxInt
	}

	q := &Queue[int]{}
	q.Enqueue(k)
	distances[k] = 0
	for q.Len() > 0 {
		v := q.Dequeue()
		for _, adjV := range adj[v] {
			if distances[v]+adjV[1] < distances[adjV[0]] {
				q.Enqueue(adjV[0])
				distances[adjV[0]] = distances[v] + adjV[1]
			}
		}
	}

	maxDistance := 0
	for _, v := range distances {
		if v == math.MaxInt {
			return -1
		} else {
			maxDistance = max(maxDistance, v)
		}
	}

	return maxDistance
}

// func networkDelayTimeDFS(times [][]int, n int, k int) int {
// 	const MAX_DISTANCE = 101
// 	adj := make([][][2]int, n+1)
// 	for _, v := range times {
// 		adj[v[0]] = append(adj[v[0]], [2]int{v[1], v[2]})
// 	}
// 	distances := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		distances[i] = MAX_DISTANCE
// 	}

// 	var dfs func(v, d int)
// 	dfs = func(v, d int) {
// 		if distances[v] > d {
// 			distances[v] = d

// 			for _, adjV := range adj[v] {
// 				dfs(adjV[0], d+adjV[1])
// 			}
// 		}
// 	}
// 	dfs(k, 0)

// 	maxDistance := 0
// 	for _, d := range distances {
// 		if d == MAX_DISTANCE {
// 			return -1
// 		} else {
// 			maxDistance = max(maxDistance, d)
// 		}
// 	}

// 	return maxDistance
// }
