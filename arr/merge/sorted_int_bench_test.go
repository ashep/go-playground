package merge_test

import (
	"go-playground/arr/merge"
	"testing"
)

func BenchmarkSortedIntV1(b *testing.B) {
	benchmarkSortedInt(b, merge.SortedIntV1)

}

func BenchmarkSortedIntV2(b *testing.B) {
	benchmarkSortedInt(b, merge.SortedIntV2)

}

func benchmarkSortedInt(bb *testing.B, fn func([]int, []int) []int) {
	var a1Of1K = makeArr(1000)
	var a2Of1K = makeArr(1000)
	var a1Of10K = makeArr(10000)
	var a2Of10K = makeArr(10000)
	var a1Of100K = makeArr(100000)
	var a2Of100K = makeArr(1000000)
	var a1Of1M = makeArr(1000000)
	var a2Of1M = makeArr(1000000)

	bb.Run("1K", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fn(a1Of1K, a2Of1K)
		}
	})

	bb.Run("10K", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fn(a1Of10K, a2Of10K)
		}
	})

	bb.Run("100K", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fn(a1Of100K, a2Of100K)
		}
	})

	bb.Run("1M", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fn(a1Of1M, a2Of1M)
		}
	})
}

func makeArr(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}

	return r
}
