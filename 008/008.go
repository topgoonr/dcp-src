/*

This problem was asked by Google.

A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.

For example, the following tree has 5 unival subtrees:

   0
  / \
 1   0
    / \
   1   0
  / \
 1   1

*/

package main

import (
	"fmt"

	Tr "../treeutil"
)

// NOTUNIVAL is an undesirable return value
const NOTUNIVAL Tr.ValueType = -1

// CheckUnivalTree returns the value of the tree
func CheckUnivalTree(t *Tr.TreeNode, someval Tr.ValueType, counter *int) Tr.ValueType {
	if t.Left == nil && t.Right == nil {
		fmt.Println("some poor leaf node", t.Value)
		*counter++
		return t.Value
	}

	// induction part
	var leftVal, rightVal Tr.ValueType
	if t.Left != nil {
		leftVal = CheckUnivalTree(t.Left, t.Value, counter)
	}
	if t.Right != nil {
		rightVal = CheckUnivalTree(t.Right, t.Value, counter)
	}

	if leftVal != t.Value {
		return NOTUNIVAL
	}
	if rightVal != t.Value {
		return NOTUNIVAL
	}
	fmt.Println("does it come here?", t)
	*counter++
	return t.Value // all subtrees have the same value as this one
}

func main() {
	instructionsString := "0 root,1 al,2 ar,7 tr al,3 tr ar,7 tr tl al,7 tr tl ar"
	mytree := Tr.CreateTree(instructionsString)
	if mytree == nil {
		fmt.Println("why?")
	}

	fmt.Println("mytree = ", mytree)
	Tr.Preorder(mytree)
	mycounter := 0

	// has to be done for every interior node
	CheckUnivalTree(mytree, mytree.Value, &mycounter)

	fmt.Println("unival trees = ", mycounter)
}
