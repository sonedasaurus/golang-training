package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var before, after []string
	var sc = bufio.NewScanner(fp)
	for sc.Scan() {
		s := expand(sc.Text(), f)
		before = append(before, sc.Text())
		after = append(after, s)
	}
	fmt.Println("-----------------\nbefore:")
	for _, value := range before {
		fmt.Println(value)
	}
	fmt.Println("-----------------\nafter:")
	for _, value := range after {
		fmt.Println(value)
	}

}

func expand(s string, f func(string) string) string {
	fields := strings.Fields(s)
	for _, field := range fields {
		if strings.HasPrefix(field, "$") {
			word := strings.TrimLeft(field, "$")
			s = strings.Replace(s, word, f(word), 1)
		}
	}
	return s
}

func f(s string) string {
	if s == "foo" {
		return "foo_text"
	}
	if s == "bar" {
		return "bar_text"
	}
	if s == "hoge" {
		return "hoge_text"
	}
	return s
}
