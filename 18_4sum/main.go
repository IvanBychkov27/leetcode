/* https://leetcode.com/problems/4sum/

Учитывая массив чисел из n целых чисел, верните массив всех уникальных квадруплетов [числа [a], числа[b], числа[c], числа[d]] таких, что:
0 <= a, b, c, d < n
a, b, c и d различны.
числа[a] + числа[b] + числа[c] + числа[d]== цель
Вы можете вернуть ответ в любом порядке.

Выполнено:
Время выполнения: 240 мс, быстрее, чем 14,35 % отправленных онлайн-сообщений для 4Sum.
Использование памяти: 6,7 МБ, менее 21,30 % онлайн-заявок на 4sum.

*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//nums, target := []int{1, 0, -1, 0, -2, 2}, 0 // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	//nums, target := []int{2, 2, 2, 2, 2}, 8 // Output: [[2,2,2,2]]
	nums, target := []int{1, -2, -5, -4, -3, 3, 3, 5}, -11 // Output: [[-5,-4,-3,1]]

	res := fourSum(nums, target)
	fmt.Println("res =", res)
}

func fourSum(nums []int, target int) [][]int {
	data := make(map[string][]int)

	funcRes := func(data map[string][]int) [][]int {
		res := make([][]int, 0)
		for _, val := range data {
			res = append(res, val)
		}
		return res
	}

	sort.Ints(nums)

	l := len(nums)
	if l == 0 {
		return funcRes(data)
	}

	pred := 1
	for i := 0; i < l-3; i++ {
		if nums[i] == pred {
			continue
		}
		if target >= 0 && nums[i] > target {
			return funcRes(data)
		}
		pred = nums[i]

		for j := i + 1; j < l-2; j++ {
			if target >= 0 && nums[i]+nums[j] > target {
				break
			}

			for k := j + 1; k < l-1; k++ {
				if target >= 0 && nums[i]+nums[j]+nums[k] > target {
					break
				}

				for p := k + 1; p < l; p++ {
					sum := nums[i] + nums[j] + nums[k] + nums[p]
					if sum > target {
						break
					}

					if sum == target {
						key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[k]) + strconv.Itoa(nums[p])
						if _, ok := data[key]; ok {
							continue
						}
						data[key] = []int{nums[i], nums[j], nums[k], nums[p]}
					}
				}
			}
		}
	}

	return funcRes(data)
}
