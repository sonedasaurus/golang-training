package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := []byte("テスト \t \n \v \f \r テスト")
	fmt.Println(string(dupRemove(a)))
}

func dupRemove(s []byte) []byte {
	count := 0
	for i := 0; i < len(s)-(count+1); i = i + 1 {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			s[i] = 0x20 // ASCII Space
			copy(s[i:], s[i+1:])
			i -= 1
			count += 1
		}
	}
	return s[:len(s)-count]
}
