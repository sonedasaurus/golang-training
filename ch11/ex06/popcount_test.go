package popcount_test

import (
	"testing"

	"./"
)

func benchmark(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B)  { benchmark(b, popcount.PopCount) }
func BenchmarkPopCount2(b *testing.B) { benchmark(b, popcount.PopCount2) }
func BenchmarkPopCount3(b *testing.B) { benchmark(b, popcount.PopCount3) }
func BenchmarkPopCount4(b *testing.B) { benchmark(b, popcount.PopCount4) }
