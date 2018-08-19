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

func check(list []int, mytarget int) bool {
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
	fmt.Println(check(givenList, target))
}
