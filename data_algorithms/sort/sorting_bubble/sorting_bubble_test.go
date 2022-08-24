package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortingByChoice(t *testing.T) {
	a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0, -2, -1, -3, -5, -4, -6, -7, -8, -9, -10}
	res := sortingBubble(a)

	expected := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, expected, res)
}

func Benchmark_sortingByChoice(b *testing.B) {
	a := setData(30000)
	for i := 0; i < b.N; i++ {
		sortingBubble(a)
	}
}
