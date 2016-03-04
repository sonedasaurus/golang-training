package main

import "strings"

func echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(args, " ")
}
