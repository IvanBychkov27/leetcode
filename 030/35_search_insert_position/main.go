/* https://leetcode.com/problems/search-insert-position/
Учитывая отсортированный массив различных целых чисел и целевое значение, верните индекс, если целевое значение найдено.
Если нет, верните индекс туда, где он был бы, если бы он был вставлен по порядку.
Вы должны написать алгоритм со сложностью выполнения O(log n).

Пример 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Пример 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Пример 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Выполнено:
Время выполнения: 8 мс, быстрее, чем 25,92% онлайн-заявок на поиск позиции вставки.
Использование памяти: 3 МБ, менее 8,44% отправленных онлайн-заявок на поиск позиции вставки.

*/
package main

import (
	"fmt"
)

func main() {
	//nums, target := []int{1, 3, 5, 6}, 4
	nums, target := []int{1, 2, 4, 6, 7}, 3

	fmt.Println("nums   =", nums)
	fmt.Println("target =", target)
	res := searchInsert(nums, target)
	fmt.Println("result =", res)
}

func searchInsert(nums []int, target int) int {
	lowKey := 0              // первый индекс
	highKey := len(nums) - 1 // последний индекс
	if nums[lowKey] > target {
		return 0
	}
	if nums[highKey] < target {
		return highKey + 1
	}
	var idx int
	for lowKey <= highKey {
		idx = (lowKey + highKey) / 2 // середина
		if nums[idx] == target {     // мы нашли значение
			return idx
		}
		if nums[idx] < target { // если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = idx + 1
			continue
		}
		if nums[idx] > target { // если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = idx - 1
		}
	}

	if nums[idx] < target {
		idx++
	}

	return idx
}
