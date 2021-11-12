// https://leetcode.com/problems/two-sum/
package main

import (
	"fmt"
)

/*
Учитывая массив целых чисел nums и целочисленную цель, возвращайте индексы двух чисел так, чтобы они складывались в цель.
Вы можете предположить, что каждый вход будет иметь ровно одно решение, и вы не можете использовать один и тот же элемент дважды.
Вы можете вернуть ответ в любом порядке.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println("res1 =", twoSum(nums, target))
	fmt.Println("res2 =", twoSum2(nums, target))
	fmt.Println("res3 =", twoSum3(nums, target))
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	r := make(map[int]int)
	for i, v := range nums {
		idx, ok := r[target-v]
		if ok {
			return []int{idx, i}
		}
		r[v] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	var i int
	l := len(nums)
	r := make(map[int]int)
	for {
		idx, ok := r[target-nums[i]]
		if ok {
			return []int{idx, i}
		}
		r[nums[i]] = i
		i++

		if i == l {
			break
		}
	}
	return nil
}
