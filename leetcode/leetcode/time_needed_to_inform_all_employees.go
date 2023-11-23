package leetcode

import "fmt"

func NumOfMinutes() {
	{
		n := 8
		headID := 4
		manager := []int{2, 2, 4, 6, -1, 4, 4, 5}
		informTime := []int{0, 0, 4, 0, 7, 3, 6, 0}
		fmt.Println(numOfMinutes(n, headID, manager, informTime))
	}

	{
		n := 1
		headID := 0
		manager := []int{-1}
		informTime := []int{0}
		fmt.Println(numOfMinutes(n, headID, manager, informTime))
	}

	{
		n := 6
		headID := 2
		manager := []int{2, 2, -1, 2, 2, 2}
		informTime := []int{0, 0, 1, 0, 0, 0}
		fmt.Println(numOfMinutes(n, headID, manager, informTime))
	}

	{
		n := 11
		headID := 4
		manager := []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}
		informTime := []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}
		fmt.Println(numOfMinutes(n, headID, manager, informTime))
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adj := make([][]int, n)
	for e, m := range manager {
		if m != -1 {
			adj[m] = append(adj[m], e)
		}
	}

	var travelDFS func(manager int) int
	travelDFS = func(manager int) int {
		if len(adj[manager]) == 0 {
			return 0
		}
		minutes := 0
		for _, v := range adj[manager] {
			minutes = max(minutes, travelDFS(v))
		}
		return minutes + informTime[manager]
	}
	return travelDFS(headID)
}
