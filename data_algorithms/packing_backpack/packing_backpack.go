/* https://www.youtube.com/watch?v=4E-_uzO0A_A&list=PLRDzFCPr95fL_5Xvnufpwj2uYZnZBBnsr&index=15
Задача об укладке рюкзака
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// предмет
type Item struct {
	Weight         int     // вес
	Value          int     // ценность
	SpecificWeight float64 // удельный вес
}

func main() {
	//items := []Item{
	//	{1, 9, 0},
	//	{2, 8, 0},
	//	{3, 7, 0},
	//	{4, 6, 0},
	//	{5, 5, 0},
	//	{6, 4, 0},
	//	{7, 3, 0},
	//	{8, 2, 0},
	//	{9, 1, 0},
	//}

	rand.Seed(time.Now().UnixNano())
	backpack := 100000 // максимальный вес рюкзака
	n := 100
	items := make([]Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Item{rand.Intn(backpack) + 1, rand.Intn(backpack) + 1, 0})
	}

	resultBackpackValue := packingBackpack(items, backpack)

	fmt.Println("value backpack   =", resultBackpackValue)
	fmt.Println()
	fmt.Println("greedy algorithm =", greedyAlgorithm(items, backpack))
}

// динамический алгоритм решения
func packingBackpack(items []Item, backpack int) int {
	field := setField(len(items), backpack+1)

	for sizeBackpack := 1; sizeBackpack <= backpack; sizeBackpack++ {
		for j := 0; j < len(items); j++ {
			var a, b, w int
			if j > 0 {
				a = field[sizeBackpack][j-1]
			}
			w = items[j].Weight
			if w <= sizeBackpack {
				b = items[j].Value
				if j > 0 {
					b += field[sizeBackpack-w][j-1]
				}
			}
			field[sizeBackpack][j] = max(a, b)
		}
	}

	//printField(field)
	//fmt.Println()

	return field[backpack][len(items)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func setField(items, backpack int) [][]int {
	field := make([][]int, backpack)
	for i := 0; i < backpack; i++ {
		field[i] = make([]int, items)
	}
	return field
}

func printField(field [][]int) {
	for _, d := range field {
		fmt.Println(d)
	}
}

// жадный алгоритм
func greedyAlgorithm(items []Item, backpack int) int {
	for i := 0; i < len(items); i++ {
		items[i].SpecificWeight = float64(items[i].Value) / float64(items[i].Weight)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].SpecificWeight > items[j].SpecificWeight
	})

	//fmt.Println(items)

	var w, v int
	for _, i := range items {
		if w+i.Weight > backpack {
			continue
		}
		w += i.Weight
		v += i.Value
	}

	return v
}
