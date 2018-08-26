/*
Problem 005:
Dated Aug 22, 2018
This problem was asked by Jane Street.

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair
Implement car and cdr.

JH notes:
Although, cons makes a pair, it can be used to make a list as well.
The 'car' and 'cdr' functions can then use the pairing function to add to the list.

In actual LISP, car and cdr are both fields. car points to the first element. cdr points to the rest of teh list.
*/

package main

import "fmt"

type cons struct {
	car int
	cdr *cons
}

// given function
func pair(a int, b int) cons {
	second := cons{car: b}
	first := cons{car: a, cdr: &second}
	fmt.Println(first)
	fmt.Println(second)

	return first
}

// func cons(a, b int) []int {
// 	mylist := make([]int)
// 	mylist.append(a)
// 	mylist.append(b)
// 	return mylist
// }

// TODO
func car(something cons) int {
	return something.car
}

// TODO
func cdr(something cons) *cons {
	return something.cdr
}

func main() {
	somepair := pair(1, 2)
	fmt.Println(car(somepair))
	fmt.Println(cdr(somepair))
}
