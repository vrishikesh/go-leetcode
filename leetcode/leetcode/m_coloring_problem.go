package leetcode

import "fmt"

func MColoring(N, M, E int, Edges [][2]int) {
	var f func(n int) int
	var m = map[int]int{}

	f = func(n int) int {
		if n > N {
			return 1
		}

		for c := 1; c <= M; c++ {
			if isValidColor(n, c, Edges, m) {
				m[n] = c
				if f(n+1) == 1 {
					return 1
				}
				m[n] = 0
			}
		}
		return 0
	}

	fmt.Println("M Coloring:", f(1))
	fmt.Println("Colors:", m)
}

func isValidColor(n, c int, Edges [][2]int, m map[int]int) bool {
	for _, edge := range Edges {
		if edge[0] == n {
			ec := m[edge[1]]
			if ec == c {
				return false
			}
		}

		if edge[1] == n {
			ec := m[edge[0]]
			if ec == c {
				return false
			}
		}
	}

	return true
}
