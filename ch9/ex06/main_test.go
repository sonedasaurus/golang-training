package main

import (
	"bytes"
	"testing"
)

func BenchmarkDrawParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		drawParallel(&buf, 100)
	}
}
