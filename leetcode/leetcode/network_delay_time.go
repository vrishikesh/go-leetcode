package leetcode

import (
	"fmt"
	"math"
	"slices"
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
		fmt.Printf("%d\n\n", networkDelayTime(times, n, k)) // 2
	}

	{
		times := [][]int{
			{1, 2, 1},
		}
		n := 2
		k := 1
		fmt.Printf("%d\n\n", networkDelayTime(times, n, k)) // 1
	}

	{
		times := [][]int{
			{1, 2, 1},
		}
		n := 2
		k := 2
		fmt.Printf("%d\n\n", networkDelayTime(times, n, k)) // -1
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
		fmt.Printf("%d\n\n", networkDelayTime(times, n, k)) // 14
	}
}

type NetworkDelayTimePQ struct {
	heap      []int
	distances []int
}

func (pq *NetworkDelayTimePQ) Len() int {
	return len(pq.heap)
}

func (pq *NetworkDelayTimePQ) Less(i, j int) bool {
	d := pq.distances
	return d[i] < d[j]
}

func (pq *NetworkDelayTimePQ) Swap(i, j int) {
	h := pq.heap
	h[i], h[j] = h[j], h[i]
}

func (pq *NetworkDelayTimePQ) Pop() any {
	h := pq.heap
	n := pq.Len()
	x := h[n-1]
	pq.heap = h[:n-1]
	return x
}

func (pq *NetworkDelayTimePQ) Push(x any) {
	pq.heap = append(pq.heap, x.(int))
}

// The Bellman-Ford
func networkDelayTime(times [][]int, n int, k int) int {
	distances := make([]int, n+1)
	for i := 1; i <= n; i++ {
		distances[i] = math.MaxInt
	}
	distances[k] = 0

	for i := 0; i < n-1; i++ {
		count := 0
		for _, edge := range times {
			source := edge[0]
			target := edge[1]
			weight := edge[2]

			if distances[source] < math.MaxInt &&
				distances[source]+weight < distances[target] {
				distances[target] = distances[source] + weight
				count += 1
			}
		}
		if count == 0 {
			break
		}
	}

	maxD := slices.Max(distances)
	if maxD == math.MaxInt {
		return -1
	}
	return maxD
}

// Dijkstra
// func networkDelayTime(times [][]int, n int, k int) int {
// 	adj := make([][][2]int, n+1)
// 	for _, v := range times {
// 		adj[v[0]] = append(adj[v[0]], [2]int{v[1], v[2]})
// 	}

// 	distances := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		distances[i] = math.MaxInt
// 	}

// 	pq := NetworkDelayTimePQ{[]int{}, distances}
// 	heap.Init(&pq)
// 	pq.Push(k)
// 	distances[k] = 0

// 	for pq.Len() > 0 {
// 		v := pq.Pop().(int)
// 		for _, adjV := range adj[v] {
// 			if distances[v]+adjV[1] < distances[adjV[0]] {
// 				distances[adjV[0]] = distances[v] + adjV[1]
// 				pq.Push(adjV[0])
// 			}
// 		}
// 	}

// 	maxD := slices.Max(distances)
// 	if maxD == math.MaxInt {
// 		return -1
// 	}
// 	return maxD
// }

// DFS
// func networkDelayTime(times [][]int, n int, k int) int {
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
