package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		buf.Write([]byte(s))
		return buf.String()
	}
	r := len(s) % 3
	if r != 0 {
		buf.Write([]byte(s[0:r]))
		buf.WriteByte(',')
	}
	for i := r; i < len(s); i = i + 3 {
		buf.Write([]byte(s[i : i+3]))
		if i+3 < len(s) {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
