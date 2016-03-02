package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := [32]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)

	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints [32]int
		for i, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints[i] = int(x)
		}
		reverse(&ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

func reverse(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
