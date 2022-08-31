/* см: https://habr.com/ru/post/335920/

	Быстрая сортировка

	Сложность алгоритма: худшее время - квадратичное O(n2); лучшее - линейно-логарифмическое время O(n log n)
	Идея заключается:
    - взять первое значение как опорную точку (лучше выбирать не первый, а произвольный)
    - перебирать остальные элементы
    - разделить элементы на те, которые больше и те, которые меньше опорной точки
    - рекурсивно отсортировать меньшие и большие элементы
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var count int64

func main() {
	var timeStart time.Time
	var sortA []int

	dataDisplay := false
	//dataDisplay := true

	a := setData(30000)
	//a := setDataRandom(30000)
	//a := []int{1, 1, -2, -1, 1}
	//a := []int{2, 2, 2, 2, -2}
	//a := []int{-1, -1, 1, 2, -1}
	//a := []int{1, 3, -3, 0, -1, -2, -2, -2}
	//a := []int{1, 3, -3, 0, -1, -2}
	//a := []int{-2, -4, 2, 0, -3, -3, 1, 2}
	//a := []int{1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 10, 1, 1, 1, 1}
	if dataDisplay {
		fmt.Println(a)
		fmt.Println()
	}

	a1 := make([]int, 0, len(a))
	a2 := make([]int, 0, len(a))
	a3 := make([]int, 0, len(a))
	a4 := make([]int, 0, len(a))
	for _, d := range a {
		a1 = append(a1, d)
		a2 = append(a2, d)
		a3 = append(a3, d)
		a4 = append(a4, d)
	}

	fmt.Printf("Version 1: len = %d time = ", len(a))
	timeStart = time.Now()
	sortA = quickSortV1(a1)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	if dataDisplay {
		fmt.Println(sortA)
	}
	fmt.Println("count:", count)

	fmt.Printf("\nVersion 2: len = %d time = ", len(a))
	count = 0
	timeStart = time.Now()
	sortA = quickSortV2(a2)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	if dataDisplay {
		fmt.Println(sortA)
	}
	fmt.Println("count:", count)

	fmt.Printf("\nVersion 3: len = %d time = ", len(a))
	count = 0
	timeStart = time.Now()
	sortA = quickSortV3(a3)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	if dataDisplay {
		fmt.Println(sortA)
	}
	fmt.Println("count:", count)

	if equalArrays(sortA, quickSortV2(a)) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR")
	}

	fmt.Printf("\nVersion 4: len = %d time = ", len(a))
	count = 0
	timeStart = time.Now()
	sortA = quickSortV4(a4, false)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	if dataDisplay {
		fmt.Println(sortA)
	}
	fmt.Println("count:", count)

	if equalArrays(sortA, quickSortV2(a)) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR")
	}

}

// быстрая сортировка версия 1
// Benchmark_quickSortV1-8   	       1	1614648325 ns/op	7422536856 B/op	   60036 allocs/op (тест проводился на 30 000 элементах отсортированных в обратном порядке)
// Version 1: len = 30000 time = 1.746544394s
// count: 449985000
// // O(N^2)
func quickSortV1(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	less := make([]int, 0, n)
	bigger := make([]int, 0, n)
	d := data[0] // рекомендация: данный элемент лучше выбирать не первым, а произвольным (например: d := data[idx], где idx = len(data)/2) - это чаще всего ускоряет работу функции
	for _, v := range data[1:] {
		count++
		if v > d {
			bigger = append(bigger, v) // рекомендация: для экономии памяти необходимо менять местами эл-ты в том же массиве, а не создавать новые массивы рекурсивно
		} else {
			less = append(less, v)
		}
	}
	data = append(quickSortV1(less), d)
	data = append(data, quickSortV1(bigger)...)
	return data
}

// быстрая сортировка версия 2
// Benchmark_quickSortV2-8        	     322	   3783476 ns/op	10729655 B/op	   45053 allocs/op (тест проводился на 30 000 элементах отсортированных в обратном порядке)
// Version 2: len = 30000 time = 3.610804ms
// count: 403631
// O(2*N*LogN)
func quickSortV2(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	less := make([]int, 0, n)
	bigger := make([]int, 0, n)
	idx := n / 2
	d := data[idx]
	for i, v := range data {
		count++
		if i == idx {
			continue
		}
		if v > d {
			bigger = append(bigger, v) // рекомендация: для экономии памяти необходимо менять местами эл-ты в том же массиве, а не создавать новые массивы рекурсивно
		} else {
			less = append(less, v)
		}
	}
	data = append(quickSortV2(less), d)
	data = append(data, quickSortV2(bigger)...)
	return data
}

// быстрая сортировка версия 3
// Benchmark_quickSortV3-8        	  116841	      9573 ns/op	       0 B/op	       0 allocs/op (тест проводился на 30 000 элементах отсортированных в обратном порядке)
// Version 3: len = 30000 time = 126µs
// count: 154678
// O(N*LogN) !!!
func quickSortV3(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	isSort := true
	for i := 1; i < n; i++ {
		count++
		if a[i-1] > a[i] {
			isSort = false
			break
		}
	}
	if isSort {
		return a
	}

	//idx := n / 2
	idx := rand.Intn(n)

	d := a[idx]
	i := 0
	j := n - 1

	for i < j {
		count++

		if i == idx {
			if a[j] <= d {
				a[i], a[j] = a[j], a[i]
				idx = j
				i++
			} else {
				j--
			}
			continue
		}

		if a[i] < d {
			i++
			continue
		}

		if a[j] <= d {
			a[i], a[j] = a[j], a[i]
			if j == idx {
				idx = i
			}
		}
		j--

	}

	return append(quickSortV3(a[:idx]), quickSortV3(a[idx:])...)
}

// быстрая сортировка версия 4
// Benchmark_quickSortV4-8        	  81295	     14728 ns/op	       0 B/op	       0 allocs/op (тест проводился на 30 000 элементах отсортированных в обратном порядке)
// Version 4: len = 30000 time = 16µs
// count:  59998
// O(2*N) !!!
func quickSortV4(a []int, random bool) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	isSort := true
	for i := 1; i < n; i++ {
		count++
		if a[i-1] > a[i] {
			isSort = false
			break
		}
	}
	if isSort {
		return a
	}

	idx := n / 2
	if random {
		idx = rand.Intn(n)
	}

	d := a[idx]
	i := 0
	j := n - 1

	for i < j {
		count++
		if i == idx {
			if a[j] <= d {
				a[i], a[j] = a[j], a[i]
				idx = j
				i++
			} else {
				j--
			}
			continue
		}

		if a[i] < d {
			i++
			continue
		}

		if a[j] <= d {
			a[i], a[j] = a[j], a[i]
			if j == idx {
				idx = i
			}
		}
		j--

	}

	random = false
	if idx == 0 || idx == n-1 {
		random = true
	}

	return append(quickSortV4(a[:idx], random), quickSortV4(a[idx:], random)...)
}

func setData(n int) []int {
	a := make([]int, 0, n)
	d := n / 2
	for i := 0; i < n; i++ {
		a = append(a, d-i)
	}
	return a
}

func setDataRandom(n int) []int {
	rand.Seed(time.Now().UnixMicro())
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(n)-n/2)
	}
	return a
}

func equalArrays(a, b []int) bool {
	n := len(a)
	m := len(b)
	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
