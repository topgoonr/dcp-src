# treeutil

## CreateTree

A function to parse tree building instructions, and actually build one. Allocation of TreeNodes in Go is very interesting, and needs to be treated specially. It is resolved here with the use of Nodes stored in a Slice.