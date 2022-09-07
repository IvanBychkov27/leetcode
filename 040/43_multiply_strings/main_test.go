package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_multiply(t *testing.T) {
	num1, num2 := "123", "456"
	expected := "56088"
	res := multiply(num1, num2)
	assert.Equal(t, expected, res)
}

func Benchmark_multiply(b *testing.B) {
	num1 := "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	num2 := "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		multiply(num1, num2)
	}
}
