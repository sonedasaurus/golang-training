package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	counts := make(map[string]int)
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var sc = bufio.NewScanner(fp)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		counts[sc.Text()]++
	}
	fmt.Printf("term\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
