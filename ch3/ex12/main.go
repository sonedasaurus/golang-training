package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("ex.\n$ go run main.go anagram margana\n")
		return
	}
	if checkAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("\"%s\" is anagram of \"%s\"\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("\"%s\" is NOT anagram of \"%s\"\n", os.Args[1], os.Args[2])
	}
}

func checkAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i = i + 1 {
		index := strings.Index(s2, s1[i:i+1])
		if index == -1 {
			return false
		}
		s2 = s2[0:index] + s2[index+1:len(s2)]
	}
	return true
}
