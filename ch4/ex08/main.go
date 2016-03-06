package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	letters := 0                    // count of letters
	numbers := 0                    // count of numbers
	uppers := 0                     // count of uppers
	lowers := 0                     // count of lowers
	spaces := 0                     // count of spaces
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF || r == 0x0a {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters++
		}
		if unicode.IsNumber(r) {
			numbers++
		}
		if unicode.IsUpper(r) {
			uppers++
		}
		if unicode.IsLower(r) {
			lowers++
		}
		if unicode.IsSpace(r) {
			spaces++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("letterCount\n")
	fmt.Printf("%d\n", letters)
	fmt.Printf("numberCount\n")
	fmt.Printf("%d\n", numbers)
	fmt.Printf("upperCount\n")
	fmt.Printf("%d\n", uppers)
	fmt.Printf("lowerCount\n")
	fmt.Printf("%d\n", lowers)
	fmt.Printf("spaceCount\n")
	fmt.Printf("%d\n", spaces)
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
