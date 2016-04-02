package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func main() {
	//!+main
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())             // "{1 9 144}"
	fmt.Printf("Len() = %d\n", x.Len()) // "Len() = 3"
	x.Remove(1)
	fmt.Printf("Remove(1) = %s\n", x.String()) // "Remove() = {9 144}"
	xcopy := x.Copy()
	xcopy.Add(2)
	fmt.Printf("x = %s\nxcopy = %s\n", x.String(), xcopy.String()) // "x = {2, 9, 144} xcopy = {9, 144}"
	x.Clear()
	fmt.Printf("Clear() = %s\n", x.String()) // "Clear() = {}"
}

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) && s.words[word]&(1<<bit) != 0 {
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	for i, word := range s.words {
		s.words[i] ^= word
	}
}

func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	for i, word := range s.words {
		copy.words = append(copy.words, 0)
		copy.words[i] |= word
	}
	return &copy
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
