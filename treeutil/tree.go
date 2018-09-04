package treeutil

import (
	"fmt"
	"strconv"
)

type valueType int
type treeNode struct {
	left  *treeNode
	right *treeNode
	value valueType
}

type direction int

const (
	left direction = iota
	right
)

func addNode(thisNode *treeNode, dir direction, value valueType) {
	newNode := treeNode{nil, nil, value}
	switch dir {

	case left:
		thisNode.left = &newNode
	default:
		thisNode.right = &newNode
	}
}

/*
For example, the following tree:

   0
  / \
 1   0
    / \
   1   0
  / \
 1   1

VALUE <traversal instructions> <add instruction>

0 root
1 la
0 ra
1 tr al
0 tr ar
1 tr tl al
1 tr tl ar

*/

func createTree(instructions [][]string) *treeNode {
	// parse the instructions
	var root *treeNode
	var thisNode *treeNode
	var newNode treeNode // an actual node, and not a pointer
	for lineNum, line := range instructions {
		for _, word := range line {
			temp, _ := strconv.Atoi(word) // ignore the err
			value := valueType(temp)
			if lineNum == 0 {
				newNode := treeNode{nil, nil, value}
				root = &newNode
				thisNode = root
			}
			if lineNum == 0 && word != "root" {
				return nil
			}
			switch word {
			case "tl":
				thisNode = thisNode.left
			case "tr":
				thisNode = thisNode.right
			case "al":
				newNode = treeNode{nil, nil, value}
				thisNode.left = &newNode
			case "ar":
				newNode = treeNode{nil, nil, value}
				thisNode.left = &newNode
			}
		}
	}
	return root
}

func preorder(t *treeNode, s string) {
	if t.left == nil && t.right == nil {
		fmt.Printf("%d ", int(t.value))
		preorder(t.left, s)
		preorder(t.right, s)
	}
}

// TODO: return as string as opposed to just dumping it on the console
func traversePreorder(t *treeNode) []valueType {
	return nil
}

func traverseInorder(t *treeNode) []valueType {
	return nil
}

func traversePostorder(t *treeNode) []valueType {
	return nil
}
