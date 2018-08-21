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
