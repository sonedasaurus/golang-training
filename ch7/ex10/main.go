package main

import "sort"

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
