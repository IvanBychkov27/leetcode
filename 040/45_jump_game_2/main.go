/* https://leetcode.com/problems/jump-game-ii/
45. Jump Game II
Вам дается 0-индексированный массив целых чисел nums длины n. Изначально вы находитесь в nums[0].
Каждый элемент nums[i] представляет максимальную длину прямого перехода от индекса i.
Другими словами, если вы находитесь в nums[i], вы можете перейти к любому nums[i + j], где: 0 <= j <= числа[i] и i + j < n
Возвращает минимальное количество переходов для достижения nums[n - 1]. Тестовые примеры генерируются таким образом, чтобы вы могли достичь nums[n - 1].

Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Пояснение: Минимальное количество переходов для достижения последнего индекса равно 2. Перейдите на 1 шаг от индекса 0 к 1, затем на 3 шага к последнему индексу.

Example 2:
Input: nums = [2,3,0,1,4]
Output: 2

Выполнено:
Время выполнения: 27 мс, быстрее, чем 73,71% онлайн-заявок на Jump Game II.
Объем используемой памяти: 5,9 МБ, что составляет менее 92,13% от общего количества отправленных онлайн-заявок на Jump Game II.

*/
package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{2, 3, 2, 0, 1, 9} // 3
	//nums := []int{2, 3, 1, 1, 4} // 2
	//nums := []int{2, 3, 0, 1, 4} // 2
	//nums := []int{2, 3} // 1
	//nums := []int{0} // 0
	start := time.Now()
	res := jump(nums)
	fmt.Println("time  =", time.Now().Sub(start))
	fmt.Println("res  =", res)
}

func jump(nums []int) int {
	l := len(nums) - 1
	if l == 0 {
		return 0
	}
	idx := l
	count := 0
	for idx != 0 {
		count++
		idx = idxStep(nums[:idx+1])
	}
	return count
}

func idxStep(nums []int) int {
	l := len(nums) - 1
	min := l
	target := l
	for i := 0; i < l; i++ {
		if i+nums[i] >= target && i < min {
			min = i
		}
	}
	return min
}
