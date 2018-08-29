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


1. Linux Journal has an excellent article on this. This is the seminal article that almost everyone refers to. Well written by Prokash Sinha.
https://www.linuxjournal.com/article/6828?page=0,0
2. unf.edu has a book on this. It is brief, but should work:
https://www.unf.edu/~wkloster/3540/wiki_book.pdf


The XOR 'both' field is also known as the pointer difference.
So, the 'add'
 function looks like this:
add:
	this.both = this.both ^ newnode
	newnode.both = this.both ^ nil

The node at the head always has a nil in the 'both' field
*/
package main

import (
	"fmt"
	"strconv"
)

// using indices to an array instead of pointers
type xorlist struct {
	value string
	both  int // previous XOR next --> both
}

var head int // points to the head index of the list, should be initialized to zero

// test using this:
var globalList = []xorlist{{value: "head", both: 0}}

// head is added only as a necessary filler.
// Otherwise there is no convenient way to simulate a NIL with indices

// TODO: add a new node
/* technique:
add:
	this.both = this.both ^ newnode
	newnode.both = this.both ^ nil
*/
func addToTail(previndex, thisindex int, newvalue string) int {
	fmt.Println("prev, this", previndex, thisindex)
	newBothValue := int(thisindex ^ 0)
	tempNew := xorlist{value: newvalue, both: newBothValue}
	// recalculate this.both
	fmt.Println("here2")
	newIndex := len(globalList)
	thisBothValue := int(previndex ^ newIndex)
	globalList = append(globalList, tempNew)
	// create the entry first

	globalList[thisindex].both = thisBothValue
	fmt.Println("here3")
	// 	add the new node to globalList
	return len(globalList) - 1
}

// TODO: traversal to the next node is enabled here
func nextNode(previndex, thisindex int) int {
	// XOR of this node.both and prevnode ptr
	return int(previndex ^ globalList[thisindex].both)
}

// TODO: get operation
func get(thisindex int) string {
	return globalList[thisindex].value
}

// IGNORE THIS SECTION. Attempted pointer version.

func main() {
	// TODO: Create an XOR linked list

	for i := len(globalList) - 1; i < 6; i++ {
		fmt.Println(i)
		prev := i
		somevalue := strconv.Itoa((i + 1) * 10)
		addToTail(prev, i+1, somevalue)
	}

	fmt.Println(globalList)
	for i := range globalList {
		fmt.Printf("%d = %s <%4b>\n", i, get(i), byte(globalList[i].both))
	}

	// TODO: test using a forward traversal
	// for this := 0;
	// TODO: test using a backward traversal

}
