/* https://leetcode.com/problems/3sum/
Учитывая целочисленный массив nums, верните все триплеты [nums[i], nums[j], nums[k]], такие что
  i != j, i!= k и j!= k, и
  nums[i] + nums[j] + nums[k]== 0.
Обратите внимание, что набор решений не должен содержать повторяющиеся тройки.

Решено:
Время выполнения: 2216 мс, быстрее, чем 7,62 % отправленных онлайн-сообщений для 3sum.
Использование памяти: 10 МБ, менее 8,41 % отправленных онлайн-сообщений для 3sum.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{-1, 0, 1, 2, -1, -4, -8, -7, -6, 3, 4, 5, 6, 7, 8, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	//nums := []int{}
	//nums := []int{0}

	res := threeSum(nums)

	fmt.Println("res =", res)
}

func threeSum(nums []int) [][]int {
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
	for i := 0; i < l; i++ {
		if nums[i] == pred {
			continue
		}
		if nums[i] > 0 {
			return funcRes(data)
		}
		pred = nums[i]

		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] > 0 {
				break
			}

			for k := j + 1; k < l; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum > 0 {
					break
				}

				if sum == 0 {
					key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[k])
					if _, ok := data[key]; ok {
						continue
					}
					data[key] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	return funcRes(data)
}
