/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

*/

package main

import (
	"fmt"
)

func multiplier(givenList []int) []int64 {
	numElems := len(givenList)
	var newList = make([]int64, numElems) // create a slice instead
	var mult int64                        // large size for large multiples from all the numbers

	for i := 0; i < numElems; i++ {
		mult = 1 // good to avoid multiple declarations, so declared it at the top
		for j := 0; j < numElems; j++ {
			if j == i {
				continue // next iteration please
			}
			mult = mult * int64(givenList[j])
		}
		newList[i] = mult
	}
	return newList
}

func main() {

	// var originalList = []int{1, 2, 3, 4, 5}
	var originalList = []int{3, 2, 1}
	fmt.Println(multiplier(originalList))
}
