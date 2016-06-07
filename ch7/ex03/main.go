package main

import "strconv"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	return getTreeValue("", t)
}

func getTreeValue(s string, t *tree) string {
	s = strconv.Itoa(t.value)
	if t.left != nil {
		s = s + " " + getTreeValue(s, t.left)
	}
	if t.right != nil {
		s = s + " " + getTreeValue(s, t.right)
	}
	return s
}
