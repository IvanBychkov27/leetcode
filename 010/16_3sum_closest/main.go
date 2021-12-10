/* https://leetcode.com/problems/3sum-closest/

Учитывая целочисленный массив чисел длины n и целочисленную цель, найдите три целых числа в числах, чтобы сумма была ближе всего к цели.
Верните сумму трех целых чисел.
Вы можете предположить, что каждый вход будет иметь ровно одно решение.

Решено:
Время выполнения: 39 мс, быстрее, чем 16,29 % отправленных онлайн-заявок на 3Sum Ближайший.
Использование памяти: 3,1 МБ, менее 10,54 % отправленных онлайн-сообщений для 3SUM Ближайшего.

*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//nums, target := []int{-1, 2, 1, -4}, 1
	//nums, target := []int{0, 0, 0}, 1
	//nums, target := []int{100, 200, 300, 10}, 1
	//nums, target := []int{1, 1, 1, 0}, 100
	nums, target := []int{1, 2, 4, 8, 16, 32, 64, 128}, 82 // 82

	res := threeSumClosest(nums, target)
	fmt.Println("res =", res)
}

func threeSumClosest(nums []int, target int) int {
	const coef = 1000000

	l := len(nums)
	if l < 3 {
		return 0
	}

	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[2]
	delta := coef
	pred := target + coef

	for i := 0; i < l-2; i++ {
		if nums[i] == pred {
			continue
		}
		pred = nums[i]

		for j := i + 1; j < l-1; j++ {

			for k := j + 1; k < l; k++ {
				sum := nums[i] + nums[j] + nums[k]

				if sum == target {
					return sum
				}

				m := int(math.Abs(float64(target - sum)))
				if m < delta {
					delta = m
					res = sum
				}
			}
		}
	}
	return res
}
