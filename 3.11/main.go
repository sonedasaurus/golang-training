package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	var fp string
	if s[0] == '-' {
		buf.Write([]byte(s[0:1]))
		s = s[1:len(s)]
	}
	if len(s) <= 3 {
		buf.Write([]byte(s))
		return buf.String()
	}
	i := strings.Index(s, ".")
	if i != -1 {
		fp = s[i:len(s)]
		s = s[0:i]
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
	if fp != "" {
		buf.Write([]byte(fp[0:len(fp)]))
	}
	return buf.String()
}
