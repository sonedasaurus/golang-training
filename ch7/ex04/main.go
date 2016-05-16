package main

import "io"

type XReader struct {
	s        string
	i        int64
	prevRune int
}

func (r *XReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) *XReader { return &XReader{s, 0, -1} }
