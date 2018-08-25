/*
Problem statement as of Aug 21, 2018

This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.

SOLUTION:
Do a base and offset analysis (from Computer Organization)

*/

package main

import (
	"fmt"
)

func missingpositive(myarray []int) int {
	var checklist = make(map[int]bool)
	min := 0
	max := 0
	arraysize := len(myarray)

	// linear time
	/*
	   set the min positive number in the array
	   set the max positive number in the array
	   set a hash map of positive numbers
	*/
	for i := 0; i < arraysize; i++ {
		if max < myarray[i] {
			max = myarray[i]
		}
		if min > myarray[i] {
			min = myarray[i]
		}
		c := myarray[i]
		checklist[c] = true
	}
	// create a checklist map
	fmt.Println(checklist)
	// count := (max - min) + 1
	for c := min; c <= max; c++ {
		if c <= 0 {
			continue
		}
		if checklist[c] == true {
			continue
		} else {
			return c
		}
	}

	// by default the next immediate positive number, if all numbers are accounted for
	// in this case, it would mean the largest number + 1

	return (max + 1)
}

func main() {
	// var givenArray = []int{3, 4, 4, 4, -1, 1}
	// var givenArray = []int{1, 2, 0}
	var givenArray = []int{3, 4, -1, 1, 2}

	fmt.Println(missingpositive(givenArray))
}
