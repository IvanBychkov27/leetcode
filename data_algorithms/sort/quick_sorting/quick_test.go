package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_quickSortV1(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0, -2, -1, -3, -5, -4, -6, -7, -8, -9, -10}
	res := quickSortV1(a)
	expected := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Test_quickSortV2(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0, -2, -1, -3, -5, -4, -6, -7, -8, -9, -10}
	res := quickSortV2(a)
	expected := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Test_quickSortV3(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0, -2, -1, -3, -5, -4, -6, -7, -8, -9, -10}
	res := quickSortV3(a)
	expected := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Test_quickSortV4(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0, -2, -1, -3, -5, -4, -6, -7, -8, -9, -10}
	res := quickSortV4(a, false)
	expected := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Benchmark_quickSortV1(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV1(a)
	}
}
func Benchmark_quickSortV1_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV1(a)
	}
}

func Benchmark_quickSortV2(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV2(a)
	}
}

func Benchmark_quickSortV2_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV2(a)
	}
}

func Benchmark_quickSortV3(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV3(a)
	}
}

func Benchmark_quickSortV3_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV3(a)
	}
}

func Benchmark_quickSortV4(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV4(a, false)
	}
}

func Benchmark_quickSortV4_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSortV4(a, false)
	}
}

func Benchmark_go_sort(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(a)
	}
}

func Benchmark_go_sort_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(a)
	}
}
