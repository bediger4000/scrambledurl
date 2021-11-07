package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	iter := newIterator(os.Args[1:])

	for {
		str, done := iter.Next()
		if done {
			break
		}
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

type iterator struct {
	max     int
	current []int
	items   []string
}

func newIterator(ary []string) *iterator {
	l := len(ary)

	i := &iterator{
		max:     l,
		current: make([]int, l),
		items:   ary,
	}

	return i
}

// Next returns a string of the concatenated items
// in the current order. No substring ever appears
// twice in the returned string. 2nd argument is false,
// until iteration is finished.
func (i *iterator) Next() (string, bool) {
	concatenated := ""
	done := false
	for !done && !i.Unique() {
		done = i.Increment()
	}
	if done {
		return "", done
	}
	for j := 0; j < i.max; j++ {
		concatenated += i.items[i.current[j]]
	}
	done = i.Increment()
	return concatenated, done
}

func (i *iterator) Increment() bool {
	done := false
	i.current[0]++
	if i.current[0] == i.max {
		i.current[0] = 0
		var j int
		for j = 1; j < i.max; j++ {
			i.current[j]++
			if i.current[j] == i.max {
				i.current[j] = 0
			} else {
				break
			}
		}
		if j == i.max {
			done = true
		}
	}
	return done
}

func (i *iterator) Unique() bool {
	for j := 0; j < i.max; j++ {
		for k := j + 1; k < i.max; k++ {
			if i.current[j] == i.current[k] {
				return false
			}
		}
	}
	return true
}
