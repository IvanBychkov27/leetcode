/* https://leetcode.com/problems/merge-sorted-array/
88. Merge Sorted Array - Объединить отсортированный массив

Вам даны два целочисленных массива nums1 и nums2, отсортированные в порядке неубывания, и два целых числа m и n,
представляющих количество элементов в nums1 и nums2 соответственно.
Объедините nums1 и nums2 в единый массив, отсортированный в порядке убывания.
Окончательный отсортированный массив не должен быть возвращен функцией, а вместо этого должен храниться внутри массива nums 1.
Чтобы учесть это, nums1 имеет длину m + n, где первые m элементов обозначают элементы, которые должны быть объединены,
а последние n элементов имеют значение 0 и должны игнорироваться. nums2 имеет длину n.

Example 1:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]

Example 2:
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]

Example 3:
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]

Выполнено:
Время выполнения: 5 мс, быстрее, чем 33,61% отправок в режиме онлайн для массива, отсортированного слиянием.
Использование памяти: 2,4 МБ, менее 21,15% отправленных онлайн-заявок на сортировку массива слиянием.

*/
package main

import "fmt"

func main() {
	//nums1, m, nums2, n := []int{1}, 1, []int{}, 0
	//nums1, m, nums2, n := []int{0}, 0, []int{1}, 1
	//nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	nums1, m, nums2, n := []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3
	merge(nums1, m, nums2, n)
	fmt.Println("res =", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j, idx int
	t := make([]int, m+n)
	for {
		if i == m && j == n {
			break
		}
		if i == m {
			t[idx] = nums2[j]
			j++
			idx++
			continue
		}
		if j == n {
			t[idx] = nums1[i]
			i++
			idx++
			continue
		}

		if nums1[i] <= nums2[j] {
			t[idx] = nums1[i]
			i++
		} else {
			t[idx] = nums2[j]
			j++
		}
		idx++
	}

	for i, d := range t {
		nums1[i] = d
	}
}
