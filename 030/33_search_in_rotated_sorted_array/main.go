// https://leetcode.com/problems/search-in-rotated-sorted-array/
/* Поиск в Повернутом отсортированном массиве
Существует целочисленный массив nums, отсортированный в порядке возрастания (с различными значениями).
Перед передачей в вашу функцию nums, возможно, поворачивается с неизвестным индексом поворота k (1 <= k < nums.длина) так,
чтобы результирующий массив был [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-индексированный).
Например, [0,1,2,4,5,6,7] может быть повернуто в pivotindex 3 и стать [4,5,6,7,0,1,2].
Учитывая nums массива после возможного поворота и целочисленный целевой объект,
верните индекс целевого объекта, если он находится в nums, или -1, если он не находится в nums.
Вы должны написать алгоритм с O(log n) сложностью во время выполнения.

Выполнено:
Время выполнения: 6 мс, быстрее, чем 30,37% отправленных онлайн-заявок для поиска в повернутом отсортированном массиве.
Использование памяти: 2,6 МБ, менее 70,59% отправленных в Интернет материалов для поиска в повернутом отсортированном массиве.

*/
package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{1, 3}
	//nums := []int{3, 1}
	//nums := []int{1}

	//fmt.Println("min idx:", searchMin(nums))

	res := search(nums, 0)
	fmt.Println("res =", res)
}

func search(nums []int, target int) int {
	res := -1
	if len(nums) == 0 {
		return res
	}

	idxMin := searchMin(nums)
	var idx int
	var ok bool

	if idxMin != 0 {
		idx, ok = searchInIntSlice(nums[:idxMin], target)
		if ok {
			return idx
		}
	}

	idx, ok = searchInIntSlice(nums[idxMin:], target)
	if ok {
		res = idx + idxMin
	}

	return res
}

// binary_search min элемент в Повернутом (смещенном) отсортированном массиве
func searchMin(data []int) int {
	var idxMin int
	lowKey := 0              // первый индекс
	highKey := len(data) - 1 // последний индекс

	min := data[0]
	for lowKey <= highKey {
		idx := (lowKey + highKey) / 2 // середина
		if data[idx] >= min {
			lowKey = idx + 1
			continue
		}
		if data[idx] < min {
			highKey = idx - 1
			min = data[idx]
			idxMin = idx
		}
	}

	return idxMin
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
