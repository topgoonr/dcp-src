package treeutil

import (
	"fmt"
	"strconv"
	"strings"
)

type valueType int

// TreeNode is a binary tree node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value valueType
}

type direction int

// direction Left, Right
const (
	Left direction = iota
	Right
)

func addNode(thisNode *TreeNode, dir direction, value valueType) {
	newNode := TreeNode{nil, nil, value}
	switch dir {

	case Left:
		thisNode.Left = &newNode
	default:
		thisNode.Right = &newNode
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

// CreateTree creates a binary tree from the instructions
func CreateTree(instructions string) *TreeNode {
	instructionlines := strings.Split(instructions, "\n")
	// parse the instructions
	var root *TreeNode
	var thisNode *TreeNode
	var newNode TreeNode // an actual node, and not a pointer
	for lineNum, line := range instructionlines {
		linearr := strings.Split(line, " \t")
		for _, word := range linearr {
			temp, _ := strconv.Atoi(word) // ignore the err
			value := valueType(temp)
			if lineNum == 0 {
				newNode := TreeNode{nil, nil, value}
				root = &newNode
				thisNode = root
			}
			if lineNum == 0 && word != "root" {
				return nil
			}
			switch word {
			case "tl":
				thisNode = thisNode.Left
			case "tr":
				thisNode = thisNode.Right
			case "al":
				newNode = TreeNode{nil, nil, value}
				thisNode.Left = &newNode
			case "ar":
				newNode = TreeNode{nil, nil, value}
				thisNode.Left = &newNode
			}
		}
	}
	return root
}

func preorder(t *TreeNode, s string) {
	if t.Left == nil && t.Right == nil {
		fmt.Printf("%d ", int(t.Value))
		preorder(t.Left, s)
		preorder(t.Right, s)
	}
}

// TODO: return as string as opposed to just dumping it on the console
func traversePreorder(t *TreeNode) []valueType {
	return nil
}

func traverseInorder(t *TreeNode) []valueType {
	return nil
}

func traversePostorder(t *TreeNode) []valueType {
	return nil
}