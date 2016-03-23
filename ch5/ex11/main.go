package main

import (
	"fmt"
	"os"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	//"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	var check func(items, chk []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	check = func(items, chk []string) {
		for _, item := range items {
			for _, c := range chk {
				if c == item {
					os.Exit(1)
				}
			}
			chk = append(chk, item)
			check(m[item], chk)
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
		check(m[key], nil)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
