// https://leetcode.com/problems/median-of-two-sorted-arrays/
/*
Дано два отсортированных массива nums1 и nums2 размером m и n соответственно, верните медиану двух отсортированных массивов.
Общая сложность времени выполнения должна быть O(log(m+n)).

Результат:
Время выполнения: 16 мс, быстрее, чем 60,72 % онлайн-отправлений для медианы двух отсортированных массивов.
Использование памяти: 5,8 МБ, менее 54,69 % отправленных онлайн-сообщений для медианы из двух отсортированных массивов.
*/

package main

import (
	"fmt"
)

func main() {
	//nums1 := []int{1, 2}
	//nums1 := []int{}
	//nums2 := []int{3, 4}
	//nums2 := []int{}

	nums1 := []int{1, 3}
	nums2 := []int{2}

	//nums1 := []int{0,0}
	//nums2 := []int{0,0}

	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println("res =", res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 == 0 && l2 == 0 {
		return 0
	}

	var i, j int
	nums := make([]int, 0, l1+l2)
	for {
		if i == l1 {
			nums = append(nums, nums2[j:]...)
			break
		}
		if j == l2 {
			nums = append(nums, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}

	l := len(nums)
	if l%2 == 0 {
		return (float64(nums[l/2-1]) + float64(nums[l/2])) / 2.0
	} else {
		return float64(nums[l/2])
	}
}
