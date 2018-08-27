/*

Problem 6: Aug 23, 2018
This problem was asked by Google.

An XOR linked list is a more memory efficient doubly linked list. Instead of each node holding next and prev fields, it holds a field named both, which is an XOR of the next node and the previous node. Implement an XOR linked list;
it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

If using a language that has no pointers (such as Python), you can assume you have access to get_pointer and dereference_pointer functions that converts between nodes and memory addresses.


JH Notes:
Very minimal description of an XOR list in the problem statement.

XOR linked list is a memory efficient doubly linked list.

Best explanation at this site: http://www.ritambhara.in/memory-efficient-doubly-linked-list/

Interesting properties:
The basic operations on XOR, has some interesting facts like,

if
 A ^ B = C
Then
 C ^ A = B and C ^ B = A

Golang issues:
Go does not like pointer arithmetic. Therefore, the unsafe.Pointer function has to be used for doing this. Reference:
https://stackoverflow.com/questions/47846206/bitwise-xor-on-address-in-golang


Good references are hard to find here. So, here are some. Stop looking at useless stackexchange posts and read the books which contain details for this: 
1. Linux Journal has an excellent article on this:  
https://www.linuxjournal.com/article/6828?page=0,0
2. unf.edu has a book on this. It is brief, but should work: 
https://www.unf.edu/~wkloster/3540/wiki_book.pdf



So, the 'add'
 function looks like this:
add:
	this.both = this.both ^ newnode
	newnode.both = this.both ^ nil
*/
package main

import (
	"fmt"
	"unsafe"
)

// using indices to an array instead of pointers
type xorlist struct {
	value string
	both  int // previous XOR next --> both
}

head := 0  // points to the head index of the list

// test using this:
var globalList = []xorlist{
	{"head", -1}, {"first", 0}}
// head is added only as a necessary filler. 
// Otherwise there is no convenient way to simulate a NIL with indices

// TODO: add a new node
// technique:
func add(previndex, thisindex int) int {

}

// TODO: traversal to the next node is enabled here
func nextnode(previndex, thisindex int) int {

}

// TODO: get operation
func get(previndex, thisindex int) int {
	return globalList[thisindex].value
}


// IGNORE THIS SECTION. Attempted pointer version. 

// old method of doing a node
/*
func add(thisnode *xorlist, someval int) *xorlist {
	thisNodeP := unsafe.Pointer(thisnode)
	latestBoth := (*xorlist)(unsafe.Pointer(uintptr(thisNodeP) ^ (uintptr(0))))

	// changing newnode.both
	newnode := xorlist{value: someval, both: latestBoth}

	// changing thisnode.both
	newNodeP := unsafe.Pointer(&newnode)
	tempboth := (*xorlist)(unsafe.Pointer(uintptr(newNodeP) ^ (uintptr(thisNodeP))))
	thisnode.both = tempboth

	return &newnode
}
*/

// TODO: get the value at this node
func get(thisnode *xorlist) int {
	fmt.Println(thisnode)
	return thisnode.value
}

func main() {
	// TODO: Create an XOR linked list
	// TODO: test using a forward traversal
	// TODO: test using a backward traversal

	x := xorlist{value: 1, both: nil}
	p1 := &x
	p2 := add(p1, 2)
	p3 := add(p2, 3)
	add(p3, 4)

	for t := p1; t != nil; {
		fmt.Println(get(t))
		// t.both ^ t
		temp1 := unsafe.Pointer(t.both)
		temp2 := unsafe.Pointer(t)
		t = (*xorlist)(unsafe.Pointer(uintptr(temp1) ^ uintptr(temp2)))
	}

}
