// https://www.youtube.com/watch?v=qmiLYhyMq1M
// см: https://habr.com/ru/post/335920/
/*
	Поразрядная сортировка
	Сложность алгоритма: O(n)
	Идея заключается:
		- сортируем все числа по младшему разряду
		- дале сортируем все числа по следующему разряду и т.д.
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var count int

func main() {
	//a := []int{1001, 11025, 100, 125, 27, 5, 18, 3, 96, 15, 0, 9, 12, 22, 34, 87, 70, 68, 57, 43}
	//a := []int{10, 8, 9, 7, 5, 6, 4, 3, 2, 1, 0}
	//a := []int{1, 2, 3, 0}
	//a := setData(30000)
	a := setDataRandom(20)
	//fmt.Println(a)

	timeStart := time.Now()
	res := radixSortV2(a)
	fmt.Println(time.Now().Sub(timeStart))
	//fmt.Println(res)
	fmt.Println("count =", count)
	//fmt.Println(numberOfDigits(10))

	if isSort(res) {
		fmt.Println("OK")
	} else {
		fmt.Println("No sort")
	}
}

// получение i-го разряда числа
func getDigit(number, i int) int {
	return int(float64(number)/(math.Pow(float64(10), float64(i-1)))) % 10
}

// получить максимальное количество разрядов в числе
func numberOfDigits(number int) int {
	i := 1
	n := 10
	for number >= n {
		//count++
		i++
		n *= 10
	}
	return i
}

// Поразрядная сортировка
func radixSort(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	data := a
	var res []int
	m := make(map[int][]int)

	digits := 1
	step := 10
	for digit := 1; digit < digits+1; digit++ {
		res = make([]int, 0, n)
		for _, number := range data {
			//count++

			if digit == 1 && number >= step {
				countDigit := numberOfDigits(number)
				if countDigit > digits {
					digits = countDigit
					step = int(math.Pow(float64(10), float64(digits)))
				}
			}

			d := getDigit(number, digit)
			_, ok := m[d]
			if !ok {
				m[d] = make([]int, 0, n)
			}
			m[d] = append(m[d], number)
		}

		for d := 0; d < 10; d++ {
			//count++
			v, ok := m[d]
			if !ok {
				continue
			}
			res = append(res, v...)
			//delete(m, d)
			m[d] = nil
		}
		data = res
	}
	return res
}

// Поразрядная сортировка
func radixSortV2(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	var m0, m1, m2, m3, m4, m5, m6, m7, m8, m9 []int

	digits := 1
	step := 10

	for _, number := range a {
		//count++
		if number >= step {
			countDigit := numberOfDigits(number)
			if countDigit > digits {
				digits = countDigit
				step = int(math.Pow(float64(10), float64(digits)))
			}
		}
	}

	for digit := 1; digit < digits+1; digit++ {

		m0 = make([]int, 0, n)
		m1 = make([]int, 0, n)
		m2 = make([]int, 0, n)
		m3 = make([]int, 0, n)
		m4 = make([]int, 0, n)
		m5 = make([]int, 0, n)
		m6 = make([]int, 0, n)
		m7 = make([]int, 0, n)
		m8 = make([]int, 0, n)
		m9 = make([]int, 0, n)

		for _, number := range a {
			//count++

			d := getDigit(number, digit)

			switch d {
			case 0:
				m0 = append(m0, number)
			case 1:
				m1 = append(m1, number)
			case 2:
				m2 = append(m2, number)
			case 3:
				m3 = append(m3, number)
			case 4:
				m4 = append(m4, number)
			case 5:
				m5 = append(m5, number)
			case 6:
				m6 = append(m6, number)
			case 7:
				m7 = append(m7, number)
			case 8:
				m8 = append(m8, number)
			case 9:
				m9 = append(m9, number)
			}
		}

		a = make([]int, 0, n)
		a = append(a, m0...)
		a = append(a, m1...)
		a = append(a, m2...)
		a = append(a, m3...)
		a = append(a, m4...)
		a = append(a, m5...)
		a = append(a, m6...)
		a = append(a, m7...)
		a = append(a, m8...)
		a = append(a, m9...)
	}

	return a
}

func setData(n int) []int {
	a := make([]int, 0, n)
	d := n
	for i := 0; i < n; i++ {
		a = append(a, d-i)
	}
	return a
}

func setDataRandom(n int) []int {
	rand.Seed(time.Now().UnixMicro())
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(n*100))
	}
	return a
}

func isSort(a []int) bool {
	n := len(a)
	for i := 1; i < n; i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
