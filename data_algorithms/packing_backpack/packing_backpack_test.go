package main

import (
	"math/rand"
	"testing"
	"time"
)

const (
	backpackC = 100
	nC        = 10000
)

func Benchmark_packingBackpack(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	backpack := backpackC // максимальный вес рюкзака
	n := nC
	items := make([]Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Item{rand.Intn(backpack) + 1, rand.Intn(backpack) + 1, 0})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packingBackpack(items, backpack)
	}
}

func Benchmark_greedyAlgorithm(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	backpack := backpackC // максимальный вес рюкзака
	n := nC
	items := make([]Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Item{rand.Intn(backpack) + 1, rand.Intn(backpack) + 1, 0})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greedyAlgorithm(items, backpack)
	}
}
