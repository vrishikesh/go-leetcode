package leetcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	queue := &Queue[int]{}

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
			queue.Enqueue(V)
		}
	}

	// no of test cases
	T := queue.Dequeue()

	for i := 0; i < T; i++ {
		// no of days
		N := queue.Dequeue()
		// cost of taxi
		X := queue.Dequeue()
		cost := 0
		prevD := 0

		for j := 0; j < N; j++ {
			// raining on this day?
			D := queue.Dequeue()
			if D == 1 || prevD == 1 {
				cost += X
			}
			prevD = D
		}

		fmt.Println(cost)
	}
}
