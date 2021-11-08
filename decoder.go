package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	c := make(chan string, 0)

	go permute(c, os.Args[1:])

	for str := range c {

		originalStringBytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "base64 decode: %v\n", err)
			continue
		}

		candidate := string(originalStringBytes)
		if !strings.HasPrefix(candidate, "http://") && !strings.HasPrefix(candidate, "https://") {
			continue
		}

		notPrintable := false
		for _, r := range candidate {
			if !unicode.IsPrint(r) {
				notPrintable = true
				break
			}
		}
		if notPrintable {
			continue
		}
		fmt.Println(candidate)
	}
}

// permute creates all permutations of the slice-of-strings it
// receives as a formal argument, one string by concatenating
// a permutation of the argument ary.
// The strings created by concatenating a permutation get written
// to a channel. This function closes the channel when func realPermute
// returns, so that the channel reader can just range over reading the
// channel.
func permute(c chan string, ary []string) {
	realPermute(c, "", ary)
	close(c)
}

// realPermute receives a channel (to which it writes "output" strings)
// a string produced by conatenation so far, and a slice of strings.
func realPermute(c chan string, sofar string, ary []string) {
	// recurson base case, write the entire string created by permutation,
	// to the output channel.
	if len(ary) == 1 {
		c <- sofar + ary[0]
		return
	}

	// This function will leave out 1 string in ary, and pass the remaining
	// pieces of ary as another slice to a recursive call to realPermute.
	newary := make([]string, len(ary)-1)

	for idx, str := range ary {
		var j int
		for i := 0; i < len(ary); i++ {
			if i == idx {
				continue
			}
			newary[j] = ary[i]
			j++
		}
		realPermute(c, sofar+str, newary)
	}
}
