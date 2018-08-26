/*

Problem 6: Aug 23, 2018
This problem was asked by Google.

An XOR linked list is a more memory efficient doubly linked list. Instead of each node holding next and prev fields, it holds a field named both, which is an XOR of the next node and the previous node. Implement an XOR linked list;
it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

If using a language that has no pointers (such as Python), you can assume you have access to get_pointer and dereference_pointer functions that converts between nodes and memory addresses.


JH Notes:
Very minimal description of an XOR list in the problem statement.

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

*/
package main

import (
	"fmt"
	"unsafe"
)

type xorlist struct {
	value int
	both  *xorlist // previous XOR next --> both
}

// TODO: add a new node
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

// TODO: get the value at this node
func get(thisnode *xorlist) int {
	fmt.Println(thisnode)
	return thisnode.value
}

func main() {

}
