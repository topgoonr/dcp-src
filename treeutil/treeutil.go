package treeutil

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// ValueType is an int
type ValueType int

// TreeNode is a binary tree node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value ValueType
}

type direction int

// direction Left, Right
const (
	Left direction = iota
	Right
)

func addNode(thisNode *TreeNode, dir direction, value ValueType) {
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

0 root, 1 al, 0 ar, 1 tr al, 0 tr ar, 1 tr tl al, 1 tr tl ar
*/

// CreateTree creates a binary tree from the instructions
func CreateTree(instructions string) *TreeNode {
	instructionlines := strings.Split(instructions, ",")
	fmt.Println(len(instructionlines), instructionlines)
	// parse the instructions
	var root *TreeNode
	var thisNode *TreeNode
	var newNode TreeNode // an actual node, and not a pointer
	var value ValueType

	for lineNum, line := range instructionlines {
		wordarr := strings.Split(line, " ")
		thisNode = root
		fmt.Println("----------NEW LINE words = ", wordarr, "lineNum = ", lineNum, "len =", len(wordarr), "root ")
		for i, word := range wordarr {
			// fmt.Println(i, word)
			fmt.Println("begin... this = ", thisNode)
			temp, err := strconv.Atoi(word)
			if err == nil { // no err then this is a value
				value = ValueType(temp)
				fmt.Println(i, "v>>", value)
			}
			// err != nil means it is not a value, so err need not be handled

			if lineNum == 0 && i == 0 {
				newNode = TreeNode{Left: nil, Right: nil, Value: value}
				root = &newNode
				thisNode = (*TreeNode)(unsafe.Pointer(root))
				fmt.Println("[ root created ] ")
				continue //next word, please
			}

			if thisNode == nil {
				fmt.Println("yeah, why?")
			}
			switch word {
			case "tl":
				fmt.Println(">tl...")
				fmt.Println("this = ", thisNode)
				thisNode = thisNode.Left
				fmt.Println("this = ", thisNode)
			case "tr":
				fmt.Println(">tr...")
				fmt.Println("this = ", thisNode)
				thisNode = thisNode.Right
				fmt.Println("this = ", thisNode)
			case "al":
				fmt.Println(">al...")
				// memory allocation causing issues here.
				// TODO: fix it!
				newNode = TreeNode{Left: nil, Right: nil, Value: value}
				fmt.Println("value = ", newNode.Value)
				fmt.Println("this = ", thisNode)
				thisNode.Left = &newNode
				fmt.Println("this = ", thisNode, "...al")
			case "ar":
				fmt.Println(">ar")
				// memory allocation causing issues here. As the root node gets changed, somehow
				// TODO: Fix it!
				newNode = TreeNode{Left: nil, Right: nil, Value: value}
				fmt.Println("value = ", newNode.Value)
				thisNode.Right = &newNode
				fmt.Println(thisNode, "..ar")
			default:
				continue
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
func traversePreorder(t *TreeNode) []ValueType {
	return nil
}

func traverseInorder(t *TreeNode) []ValueType {
	return nil
}

func traversePostorder(t *TreeNode) []ValueType {
	return nil
}
