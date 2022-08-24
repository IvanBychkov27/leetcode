package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchSubstringKMP(t *testing.T) {
	a := "teqwertyuiop asdfghjklzxcvbnmqwertyuiopasdfghjk gjgklzxcvbnmn asdfgqwerafsghd test"
	sub := "st"

	idx, ok := searchSubstringBMH(a, sub)

	assert.Equal(t, true, ok)
	assert.Equal(t, 80, idx)
}

func Benchmark_searchSubstringKMP(b *testing.B) {
	a := "teqwertyuiop asdfghjklzxcvbnmqwertyuiopasdfghjk gjgklzxcvbnmn asdfgqwerafsghd test"
	sub := "st"

	for i := 0; i < b.N; i++ {
		searchSubstringBMH(a, sub)
	}
}
