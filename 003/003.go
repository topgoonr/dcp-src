/*

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'

*/

// This is not exactly a solution, but a way to implement something similar
// the assert can come later
package main

import (
	"fmt"
)

type node struct {
	// Node type of a tree
	val   string
	left  *node
	right *node
}

// just a function, not a method
func create(v string, l *node, r *node) *node {
	newNode := node{v, l, r}
	return &newNode
}

func printtree(n *node) {
	if n == nil {
		return
	}
	fmt.Print("(", n.val)
	left := n.left
	right := n.right
	printtree(left)
	printtree(right)
	fmt.Print(")")

	// fmt.Println("(", printtree(left), printtree(right), ")")
	// does not work with Println, because it tries to convert it to a value
}

func main() {

	newroot := create("root", create("left", create("left.left", nil, nil), nil), create("right", nil, nil))
	printtree(newroot)
	fmt.Println()
}
