/* This problem was asked by Twitter.

Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.

For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].

Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.
*/

package main

import (
	"fmt"
	"io/ioutil"
)

func readDictionary([]string) int {
	// create the data structure here
	// returns count of strings read into the dictionary
	count := 0
	return count
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var trie [][]byte
	// [N, A] where N stands for the starting level and A stands for the branch
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))
}
