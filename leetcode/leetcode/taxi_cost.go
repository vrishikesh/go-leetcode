package leetcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type taxi_q[T comparable] []T

func (q *taxi_q[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *taxi_q[T]) Pop() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func TaxiCost() {
	/*
		2
		1 57
		0
		2 47
		1
		0
	*/
	taxiCost()
}

func taxiCost() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := &taxi_q[int]{}

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		for _, v := range strings.Fields(input) {
			V, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("could not convert string to num: V = %s\n", v)
				return
			}
			queue.Push(V)
		}
	}

	// no of test cases
	T := queue.Pop()

	for i := 0; i < T; i++ {
		// no of days
		N := queue.Pop()
		// cost of taxi
		X := queue.Pop()
		cost := 0
		prevD := 0

		for j := 0; j < N; j++ {
			// raining on this day?
			D := queue.Pop()
			if D == 1 || prevD == 1 {
				cost += X
			}
			prevD = D
		}

		fmt.Println(cost)
	}
}
