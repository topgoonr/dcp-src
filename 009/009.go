/*
This problem was asked by Airbnb.

Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5]should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?


First attempt: Try doing the shell sort equivalent here.
*/

package main

import "fmt"

//  a recursive method to do dynamic programming
func largestNonAdjacentSum(input []int) int {
	// standard base cases
	if len(input) <= 0 {
		return 0
	}

	if len(input) == 2 {
		if input[0] > input[1] {
			return input[0]
		}
		return input[1]
	}

	if len(input) == 1 {
		return input[0]
	}

	// inductive cases
	sumA := largestNonAdjacentSum(input[2:])
	sumB := largestNonAdjacentSum(input[3:])

	if sumA > sumB {
		return input[0] + sumA
	}
	return sumB + input[0]

}

func main() {

	// input := []int{5, 1, 1, 5} //
	input := []int{2, 4, 6, 2, 5}

	max := largestNonAdjacentSum(input)
	fmt.Println("sum = ", max)
}
