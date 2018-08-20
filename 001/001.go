package main

import "fmt"

/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

// create a recursive subtract largest and check function
//  Here is the solution with Brute force
// The solution picks one element and goes through every other element to check for a sum total

var target = 10
var givenList = []int{11, 3, 5, 7}

func bruteCheck(list []int, mytarget int) bool {
	// brute force
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if (list[i] + list[j]) == target {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(bruteCheck(givenList, target))

	fmt.Println(optimizedCheck(givenList, target))
}

func optimizedCheck(list []int, mytarget int) bool {
	// TODO:
	// pre-requisite: list must be sorted
	// traverse the 'i' index from left to right
	// traverse the 'j' index from right to left

	// in every  iteration: check sum of i-th element and the j-th element
	i := 0
	// listlen := len(list)
	for j := len(list) - 1; i < j; j-- {
		fmt.Println(i, j)
		if (list[i] + list[j]) == mytarget {
			return true
		}
		i++
	}
	return false
}
