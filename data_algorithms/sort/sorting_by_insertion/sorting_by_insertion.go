/* https://www.youtube.com/watch?v=jMWvNTp_wFA&list=PLA0M1Bcd0w8yF0PO0eJ9v8VlsYEowmsnJ&index=9
 	Сортировка вставкой
	Сложность: O(n^2)
Примечание:
Сортировка вставками (Сложность: O (N^2)). Один из самых простых алгоритмов сортировки - сортировка вставками (англ. insertion sort).
Его идея заключается в поддерживании в левой части массива отсортированного отрезка и постепенном его расширении вправо за счёт
перемещения очередного соседнего элемента в соответствующую ему позицию.

Алгоритм удобно использовать если к уже отсортированному массиву нужно добавить несколько элементов

*/
package main

import (
	"fmt"
	"time"
)

func main() {
	a := setData(10)
	fmt.Println(a)

	timeStart := time.Now()
	sortA := sortingByInsertion(a)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	fmt.Println(sortA)

	b := []int{-10, -11, -12}
	res := addElements(a, b)
	fmt.Println(res)
}

// sortingByInsertion - сортировка алгоритмом вставки
func sortingByInsertion(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

	return a
}

// addElements - к отсортированному массиву добавляем несколько элементов и сразу сортируем
func addElements(a, b []int) []int {
	idx := len(a)
	a = append(a, b...)
	n := len(a)

	for i := idx; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

	return a
}

func setData(n int) []int {
	a := make([]int, 0, n)
	d := n / 2
	for i := 0; i < n; i++ {
		a = append(a, d-i)
	}
	return a
}
