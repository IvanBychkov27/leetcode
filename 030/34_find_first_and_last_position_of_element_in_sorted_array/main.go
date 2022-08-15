// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
/* Найти первую и последнюю позицию элемента в отсортированном массиве
Учитывая массив целых чисел, отсортированных в порядке неубывания, найдите начальную и конечную позиции заданного целевого значения.
Если цель не найдена в массиве, верните [-1, -1].
Вы должны написать алгоритм с O(log n) сложностью во время выполнения.
Пример:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Выполнено:
Время выполнения: 19 мс, быстрее, чем 8,62% онлайн-заявок для поиска Первой и последней позиции элемента в отсортированном массиве.
Использование памяти: 4 МБ, менее 17,82% отправленных онлайн-заявок для поиска Первой и последней позиции элемента в отсортированном массиве.

*/

package main

import "fmt"

func main() {
	nums, target := []int{8, 8, 8, 8, 8}, 8
	//nums, target := []int{8, 9}, 9
	//nums, target := []int{5, 7, 7, 8, 8, 10}, 6
	//nums, target := []int{}, 0

	res := searchRange(nums, target)
	fmt.Println("res =", res)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if nums[0] == target && nums[len(nums)-1] == target {
		return []int{0, len(nums) - 1}
	}

	res := []int{}
	idx, ok := searchInIntSlice(nums, target)
	if !ok {
		return []int{-1, -1}
	}

	lowIdx, highIdx := idx, idx
	for target == nums[idx] {
		idx--
		if idx < 0 {
			break
		}
		if target == nums[idx] {
			lowIdx = idx
		}
	}
	res = append(res, lowIdx)

	idx = highIdx
	for target == nums[idx] {
		idx++
		if idx >= len(nums) {
			break
		}
		if target == nums[idx] {
			highIdx = idx
		}
	}
	res = append(res, highIdx)

	return res
}

// binary_search
func searchInIntSlice(data []int, n int) (idx int, res bool) {
	lowKey := 0              // первый индекс
	highKey := len(data) - 1 // последний индекс
	if (data[lowKey] > n) || (data[highKey] < n) {
		return // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		// уменьшаем список рекурсивно
		idx = (lowKey + highKey) / 2 // середина
		if data[idx] == n {
			res = true // мы нашли значение
			return
		}
		if data[idx] < n {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = idx + 1
			continue
		}
		if data[idx] > n {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = idx - 1
		}
	}
	return
}
