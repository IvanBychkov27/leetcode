package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_radixSort(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0}
	res := radixSort(a)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Test_radixSortV2(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0}
	res := radixSortV2(a)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, res)
}

func Benchmark_radixSort(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixSort(a)
	}
}
func Benchmark_radixSort_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixSort(a)
	}
}

func Benchmark_radixSortV2(b *testing.B) {
	a := setData(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixSortV2(a)
	}
}

func Benchmark_radixSortV2_rand(b *testing.B) {
	a := setDataRandom(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixSortV2(a)
	}
}
