package main

import "testing"

func Benchmark_threeSum(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4, -8, -7, -6, 3, 4, 5, 6, 7, 8, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		threeSum(nums)
	}
}
