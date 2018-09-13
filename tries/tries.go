package tries

// create an array of maps
// each map in turn has keys equalling the letters of the alphabet

// A trie [N, A] where N stands for the starting level and A stands for the branch

// TreeHeight starts at 0 and is as deep as the max string length found in the document
type TreeHeight int

// TrieMap is a map having approx. 26 keys for English (ideally). Crucial part of a Trie
type TrieMap map[byte]int // takes a letter and reveals the level, defaults to 0 otherwise

func addTrieBranch(lev TreeHeight, branch byte, tr []TrieMap) {
	tr[lev][branch]
}

// Trie is a crucial top level structure of a Trie
// Trie := make(map[TreeHeight]TrieMap)
