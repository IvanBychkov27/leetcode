// https://leetcode.com/problems/permutations/
/*
	Учитывая массив nums различных целых чисел, верните все возможные перестановки. Вы можете вернуть ответ в любом порядке.

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00% онлайн-заявок на перестановки.
Использование памяти: 2,7 МБ, менее 68,79% отправленных онлайн заявок на перестановки.

*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
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
		dd = append(dd, a)
		return
	}

	for i := n; i < len(d); i++ {
		d[n], d[i] = d[i], d[n]
		permutationsInArray(d, n+1)
		d[n], d[i] = d[i], d[n]
	}
}
