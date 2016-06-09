package main

import "io"

type limitReader struct {
	reader io.Reader
	limit  int64
	count  int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if int64(len(p))-lr.count <= lr.limit {
		p = p[:(lr.limit - lr.count)]
	}

	n, err = lr.reader.Read(p)
	if err != nil {
		return n, err
	}

	lr.count += int64(n)
	if lr.limit <= lr.count {
		err = io.EOF
	}
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader { return &limitReader{r, n, 0} }
