package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}

	res := twoSum(nums, target)

	assert.Equal(t, expected, res)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}

	res := twoSum(nums, target)

	assert.Equal(t, expected, res)
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	expected := []int{0, 1}

	res := twoSum(nums, target)

	assert.Equal(t, expected, res)
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 9, 112, 14, 16, 20, 22, 44, 55, 66, 1, 23, 3, 5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	target := 20
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
