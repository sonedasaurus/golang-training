package main

import (
	"math/rand"
	"testing"
	"time"
)

var seed = time.Now().UTC().UnixNano()
var rng = rand.New(rand.NewSource(seed))

// IntSet
func BenchmarkIntSetAdd100000(b *testing.B)    { benchmarkIntSetAdd(b, 100000) }
func BenchmarkIntSetAdd1000000(b *testing.B)   { benchmarkIntSetAdd(b, 1000000) }
func BenchmarkIntSetAdd10000000(b *testing.B)  { benchmarkIntSetAdd(b, 10000000) }
func BenchmarkIntSetAdd100000000(b *testing.B) { benchmarkIntSetAdd(b, 100000000) }

func benchmarkIntSetAdd(b *testing.B, size int) {
	m := IntSet{}
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}

// MapIntSet
func BenchmarkMapIntSetAdd100000(b *testing.B)    { benchmarkMapIntSetAdd(b, 100000) }
func BenchmarkMapIntSetAdd1000000(b *testing.B)   { benchmarkMapIntSetAdd(b, 1000000) }
func BenchmarkMapIntSetAdd10000000(b *testing.B)  { benchmarkMapIntSetAdd(b, 10000000) }
func BenchmarkMapIntSetAdd100000000(b *testing.B) { benchmarkMapIntSetAdd(b, 100000000) }

func benchmarkMapIntSetAdd(b *testing.B, size int) {
	m := NewMapIntSet()
	for i := 0; i < b.N; i++ {
		m.Add(rng.Intn(size))
	}
}
