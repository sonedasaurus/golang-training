package main

import (
	"fmt"
	"sort"
)

type StringSort []string

func (x StringSort) Len() int           { return len(x) }
func (x StringSort) Less(i, j int) bool { return x[i] < x[j] }
func (x StringSort) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	len := s.Len()
	for i := 0; i < len/2; i++ {
		if s.Less(i, len-(i+1)) || s.Less(len-(i+1), i) {
			return false
		}
	}
	return true
}

func main() {
	var test = StringSort{"test1", "test2", "test1"}
	var test2 = StringSort{"test1", "test2", "test3"}
	var test3 = StringSort{"test1", "test2", "test3", "test2", "test1"}
	if IsPalindrome(test) {
		fmt.Println("test is Palindrome")
	} else {
		fmt.Println("test is Palindrome")
	}
	if IsPalindrome(test2) {
		fmt.Println("test2 is Palindrome")
	} else {
		fmt.Println("test2 is not Palindrome")
	}
	if IsPalindrome(test3) {
		fmt.Println("test3 is Palindrome")
	} else {
		fmt.Println("test3 is not Palindrome")
	}
}
