/*
This problem was asked by Airbnb.

Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5]should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?
*/

package main

import "fmt"

func main() {

	type maxDetails struct {
		max   int
		start int
		iter  int
	}

	w := maxDetails{0, 0, 0}
	input := []int{5, 1, 1, 5} // 2, 4, 6, 2, 5}
	size := len(input)
	w.max = 0

	for iter := 2; iter < size; iter++ {
		for start := 0; start < size; start++ {
			sum := 0
			for j := start; j < size; j += iter {
				sum += input[j]
			}
			if w.max < sum {
				w.max = sum
				w.start = start
				w.iter = iter
			}
		}
	}
	fmt.Println("sum = ", w.max)
	fmt.Println("start = ", w.start, "iter = ", w.iter)
}
