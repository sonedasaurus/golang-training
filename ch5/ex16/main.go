package main

import "fmt"

func main() {
	fmt.Println(stringJoin(", ", "test", "test2", "test3"))
}

func stringJoin(separate string, strings ...string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str
		sep = separate
	}
	return s
}
