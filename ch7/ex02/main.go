package main

import "io"

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c ByteCounter
	c.w = w
	return &(c), &(c.count)
}
