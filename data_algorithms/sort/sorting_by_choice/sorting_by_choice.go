/* https://www.youtube.com/watch?v=eI1ZuXVOmng&list=PLA0M1Bcd0w8yF0PO0eJ9v8VlsYEowmsnJ&index=8
   	Сортировка выбором
	Сложность: O(n^2)
Примечание:
В информатике сортировка выбора-это алгоритм сортировки сравнения на месте.
Он имеет временную сложность O(n^2), что делает его неэффективным в больших списках и, как правило,
хуже, чем аналогичная сортировка вставки. Сортировка выбора отличается простотой и имеет преимущества
в производительности перед более сложными алгоритмами в определенных ситуациях,
особенно когда вспомогательная память ограничена.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//a := []int{-3, 5, 0, -8, 1, 10}
	//a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0, -2, -1, -3, -5, -4, -6, -7, -8, -9, -10}
	//a := []int{1, 0, -2, -1}
	a := setData(10000)
	//fmt.Println(a)

	timeStart := time.Now()
	sortA := sortingByChoice(a)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	//fmt.Println(sortA)
}

func setData(n int) []int {
	a := make([]int, 0, n)
	d := n / 2
	for i := 0; i < n; i++ {
		a = append(a, d-i)
	}
	return a
}

// sortingByChoice - сортировка алгоритмом сортировки выбором
func sortingByChoice(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	for i := 0; i < n-1; i++ {
		m := a[i]
		p := i
		for j := i + 1; j < n; j++ {
			if m > a[j] {
				m = a[j]
				p = j
			}
		}
		if p != i {
			a[i], a[p] = a[p], a[i]
		}
	}

	return a
}
