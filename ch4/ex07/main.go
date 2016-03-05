package main

import "fmt"

func main() {
	a := []byte("abcdef")
	byteReverse(a)
	fmt.Println(string(a))

	b := []rune("あいうabcえお")
	runeReverse(b)
	fmt.Println(string(b))
}

func byteReverse(s []byte) { // Not supported Multi Byte
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func runeReverse(s []rune) { // Supported Multi Byte
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
