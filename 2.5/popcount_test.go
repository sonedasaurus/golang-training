package popcount_test

import (
	"testing"

	"./"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2Count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount3Count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount3(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount4Count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount4(0x1234567890ABCDEF)
	}
}
