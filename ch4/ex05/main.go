package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 1, 2, 2, 3}
	fmt.Println(dupRemove(a[:])) // "[0 1 2 3]"
	//!-array

	//!+slice
	s := []int{0, 1, 1, 2, 2, 3}
	fmt.Println(dupRemove(s)) // "[0 1 2 3]"
	//!-slice

	// Interactive test of dupRemove.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		fmt.Printf("%v\n", dupRemove(ints))
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dupRemove(s []int) []int {
	count := 0
	for i := 0; i < len(s)-(count+1); i = i + 1 {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			i -= 1
			count += 1
		}
	}
	return s[:len(s)-count]
}
