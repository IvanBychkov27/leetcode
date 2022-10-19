/* https://leetcode.com/problems/maximum-subarray/
53. Maximum Subarray - Максимальный подмассив

Учитывая целочисленный массив nums, найдите смежный подмассив (содержащий хотя бы одно число), который имеет наибольшую сумму, и верните его сумму.
Подмассив - это непрерывная часть массива.

Пример 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Выполнено:
Время выполнения: 178 мс, быстрее, чем 59,99% отправок в режиме онлайн для максимального подмассива.
Использование памяти: 9,4 МБ, что составляет менее 58,41% от общего количества отправленных онлайн-заявок для максимального подмассива.

*/

package main

import "fmt"

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // 6
	//nums := []int{1} // 1
	//nums := []int{5, 4, -1, 7, 8} // 23
	//nums := []int{-1} // -1
	nums := []int{-2, -3, -1} // -1
	res := maxSubArray(nums)
	fmt.Println(res)
}

// алгоритм Кадане
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	f := true
	var sum, maxSum int
	sum = -1000000
	for _, d := range nums {
		if f {
			if d < 0 {
				sum = max(sum, d)
				maxSum = sum
				continue
			} else {
				sum = 0
				f = false
			}
		}

		sum += d
		sum = max(sum, 0)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
