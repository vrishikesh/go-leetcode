package leetcode

import "fmt"

type firstNonRepeatedNumber struct {
	m map[int]int
	s []int
	i int
}

func (t *firstNonRepeatedNumber) addNumber(x int) {
	t.m[x] += 1
	t.s = append(t.s, x)
	fmt.Println("add:", x)
	fmt.Println("array:", t.s)
}

func (t *firstNonRepeatedNumber) getFirstNonRepeatedNumber() int {
	for t.i < len(t.s) {
		v := t.s[t.i]
		fmt.Println("search:", v)
		if t.m[v] == 1 {
			return v
		}
		t.i++
	}
	return -1
}

func ImplementFirstNonRepeatedNumber() {
	t := &firstNonRepeatedNumber{
		m: map[int]int{},
		s: []int{},
	}
	t.addNumber(2)
	t.addNumber(5)
	fmt.Println("first:", t.getFirstNonRepeatedNumber()) // 2
	t.addNumber(2)
	fmt.Println("first:", t.getFirstNonRepeatedNumber()) // 5
	t.addNumber(5)
	fmt.Println("first:", t.getFirstNonRepeatedNumber()) // -1
	t.addNumber(8)
	fmt.Println("first:", t.getFirstNonRepeatedNumber()) // 5
}
