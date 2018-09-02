package main

/*
This problem was asked by Facebook.

Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not allowed.
*/

import (
	"fmt"

	"../unique"
)

// TODO:
// Recursive approach. Not working. Needs to be fixed.

func decode(input string) []string {
	// returns list and error status
	globalmap := map[string]string{
		"1":  "a",
		"2":  "b",
		"3":  "c",
		"4":  "d",
		"5":  "e",
		"6":  "f",
		"7":  "g",
		"8":  "h",
		"9":  "i",
		"10": "j",
		"11": "k",
		"12": "l",
		"13": "m",
		"14": "n",
		"15": "o",
		"16": "p",
		"17": "q",
		"18": "r",
		"19": "s",
		"20": "t",
		"21": "u",
		"22": "v",
		"23": "w",
		"24": "x",
		"25": "y",
		"26": "z"}

	outputlist := make([]string, 0)
	suffixlist := make([]string, 0)

	fmt.Println("len(input):", len(input))
	if len(input) < 1 {
		return outputlist // error present
	}

	// base condition: single element
	if len(input) == 1 {
		fmt.Println("1-->", globalmap[input])
		outputlist = append(outputlist, globalmap[input])
		fmt.Printf("input = %s, output = %s\n", input, outputlist)
		return outputlist
	}

	b := []byte(input)

	// only to do a decode fn on a slice operation  input[1:].
	// Not possible in Go directly like in python

	// make a combination with the suffixes returned by the recursive call
	// variadic function to add a list to a list, and keeping the last element empty

	// 11, and 1+1
	if len(input) == 2 {
		tempjoin := globalmap[string(b[0])] + globalmap[string(b[1])] // aa
		outputlist = append(outputlist, tempjoin)
		outputlist = append(outputlist, globalmap[string(b[:2])]) // k
		fmt.Println("2-->", outputlist)
		fmt.Printf("input = %s, output = %s\n", input, outputlist)
		return outputlist
	}

	// var join1, join2 bytes.Buffer
	var join1, join2 string
	var prefix string
	if len(input) > 2 {
		plainprefix := string(b[0])
		prefix = globalmap[plainprefix] // first letter
		// prefix = plainprefix
		fmt.Println("3-->")
		fmt.Printf("part 1: there -- prefix = <%s>\n", prefix)
		suffixlist = append(suffixlist, decode(string(b[1:]))...)
		fmt.Println("part 1: ... and back again; suffixlist = ", suffixlist)
		// combine prefix with the retured suffixlist
		// fmt.Println("suffixlist = ", suffixlist)
		for _, s := range suffixlist {
			// fmt.Println(s)
			join1 = prefix
			join1 += s
			fmt.Printf("join1 = %s\n", join1)
			outputlist = append(outputlist, join1)
		}

		// reset, and for a 2 char string
		plainprefix = string(b[:2])
		prefix = globalmap[plainprefix] // first two letters
		fmt.Printf("part 2: there -- prefix = <%s>\n", prefix)
		suffixlist = make([]string, 0)
		suffixlist = append(suffixlist, decode(string(b[2:]))...)
		fmt.Println("part 2: ... and back again; suffixlist = ", suffixlist)

		for _, s := range suffixlist {
			join2 = prefix
			// fmt.Printf("pref = %s, s = %s\n", join2, s)
			join2 += s
			// join2 += globalmap[s]
			fmt.Println("join2 = ", join2)
			outputlist = append(outputlist, join2)
			outputlist = append(outputlist, join2)

		}
	}

	outputlist = unique.Strings(outputlist)
	// fmt.Println("input: ", input)
	fmt.Printf("input = %s, output = %s\n", input, outputlist)

	// using the package 'unique' as provided by Kyle Banks, and creating
	// a string slice variant of it
	// https://github.com/KyleBanks/go-kit/blob/master/unique/unique.go

	return outputlist
}

func main() {
	// mylist := decode("111")
	mylist := decode("12121")
	fmt.Println(mylist)
	fmt.Println(len(mylist))
}
