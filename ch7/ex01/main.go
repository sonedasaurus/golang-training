package main

import (
	"bufio"
	"bytes"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	count := 0
	r := bytes.NewReader(p)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		count++
	}
	*c = WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	count := 0
	r := bytes.NewReader(p)
	s := bufio.NewScanner(r)
	for s.Scan() {
		count++
	}
	*c = LineCounter(count)
	return count, nil
}
