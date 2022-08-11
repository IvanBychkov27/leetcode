// https://leetcode.com/problems/permutations/
/*
	Учитывая набор чисел, nums, которые могут содержать дубликаты, верните все возможные уникальные перестановки в любом порядке.

Выполнено:
Время выполнения: 186 мс, быстрее, чем 5,28% онлайн-заявок Go для Permutations II.
Использование памяти: 7,4 МБ, менее 6,10% отправленных онлайн-заявок на перестановки II.

*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2}
	res := permuteUnique(nums)
	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {
	dd = make([][]int, 0)
	permutationsInArray(nums, 0)
	return dd
}

var dd [][]int

func permutationsInArray(d []int, n int) {
	if n == len(d)-1 {
		a := make([]int, 0, len(d))
		for _, t := range d {
			a = append(a, t)
		}
		if isDuplicate(a) {
			return
		}
		dd = append(dd, a)
		return
	}

	for i := n; i < len(d); i++ {
		d[n], d[i] = d[i], d[n]
		permutationsInArray(d, n+1)
		d[n], d[i] = d[i], d[n]
	}
}

func isDuplicate(x []int) bool {
	for _, d := range dd {
		count := 0
		for i, el := range d {
			if x[i] != el {
				break
			}
			count++
		}
		if count == len(d) {
			return true
		}
	}
	return false
}
