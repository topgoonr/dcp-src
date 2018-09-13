/* This problem was asked by Twitter.

Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.

For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].

Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.
*/

// This program implements a simple version of autocomplete
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../unique"
)

func swap(s []string, index1 int, index2 int) {
	t := s[index1]
	s[index1] = s[index2]
	s[index2] = t
}

// lexicographic comparison in Go is done using the 'strings.Compare(a, b)'

func partition(s []string, start int, end int) int {
	down := start
	up := end
	pivot := int((start + end) / 2)
	pivotkey := s[pivot]

	swap(s, start, pivot) // save it at the 0th spot
	for down < up {
		for ; strings.Compare(s[down], pivotkey) <= 0; down++ {
			if down > end {
				break
			}
		}
		for ; strings.Compare(pivotkey, s[up]) < 0; up-- {
		}

		if down < up {
			swap(s, down, up)
		}
	}
	swap(s, start, up)
	return up
}

// inplace sorting
func chotaSort(s []string, start int, end int) {
	if start < end {

		pivot := partition(s, start, end)
		chotaSort(s, start, pivot-1)
		chotaSort(s, pivot+1, end)
	}
}

// This is the toplevel Quick sort
func qsort(s []string) {
	chotaSort(s, 0, len(s)-1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readAndSort(infile string, outfile string) []string {
	var words []string
	var lines []string
	wordsinthefile := make([]string, 0)

	dat, err := ioutil.ReadFile(infile)
	check(err)
	fmt.Println(string(dat), len(dat))
	lines = strings.Split(string(dat), "\n")
	for _, ln := range lines {
		words = strings.Split(string(ln), " ")
		for _, wd := range words {
			fmt.Println(wd)
			wordsinthefile = append(wordsinthefile, strings.ToLower(wd))
		}
	}
	fmt.Println("total words = ", len(wordsinthefile))

	// TODO: run the quick sort on this, and store them in this format
	// word1
	// word2
	// ...

	qsort(wordsinthefile)
	// fmt.Println(wordsinthefile)

	// get the unique words only
	uniqwords := unique.Strings(wordsinthefile)
	// fmt.Println(uniqwords)
	f, err := os.Create(outfile)
	var tmp string
	for _, w := range uniqwords {
		tmp = w + "\n"
		_, err = f.WriteString(tmp)
	}
	f.Sync()

	return uniqwords
}

// use the file and return the max suggestions that start with 'input'
func listcandidates(input string, sortedwords []string, maxsuggestions int) []string {

	var wordbytes []byte
	pref := len(input) // prefix length to be considered
	// do binary search
	found := false
	start := 0
	end := len(sortedwords)
	var k int
	for k = (start + end) / 2; !strings.HasPrefix(sortedwords[k], input); k = (start + end) / 2 {
		fmt.Println("k = ", sortedwords[k], "; q = ", input, "; (", start, ",", end, ")")
		wordbytes = []byte(sortedwords[k])
		fmt.Println("? ", string(wordbytes[:pref]))
		comparison := strings.Compare(string(wordbytes[:pref]), input)
		switch {
		case comparison > 0:
			end = k
		case comparison < 0:
			start = k
		case comparison == 0:
			found = true
			fmt.Println("found!")
		}
	}

	fmt.Println("sortedwords[k] = ", sortedwords[k], ", ", "input = ", input)
	if strings.HasPrefix(sortedwords[k], input) {
		found = true
		fmt.Println("X")
	}

	fmt.Println("k =", k, "size = ", len(sortedwords), sortedwords[k])

	if !found {
		return nil
	}

	// now 'k' returns only one of the indices. You will have to traverse all the way up to check the earliest occurence of the substring
	// found the string
	// choose start and go right up to 'k'
	var i = k
	for ; strings.HasPrefix(sortedwords[i], input); i-- {
	}
	fmt.Println("i = ", i)

	firstindex := i
	lastindex := firstindex + maxsuggestions

	candidates := sortedwords[firstindex:lastindex]
	fmt.Println("firstindex = ", firstindex, "; lastindex = ", lastindex, "; size = ", len(candidates))
	return candidates
}

func main() {
	// testquery := "acc"
	uw := readAndSort("./test/2.test", "./test/2.sorted")
	fmt.Println(uw)
	suggestions := listcandidates("ye", uw, 10)
	for _, w := range suggestions {
		fmt.Printf("<%s>\n", w)
	}
	fmt.Println(suggestions)
	// started working when the overall memory footprint was reduced

	// readAndSort("./test/3.test", "./test/3.sorted")
	// readAndSort("./test/1.test", "./test/1.sorted")
}
