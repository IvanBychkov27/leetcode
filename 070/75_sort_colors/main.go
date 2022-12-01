// https://leetcode.com/problems/sort-colors/description/
/* 75. Sort Colors
Учитывая массив nums с n объектами, окрашенными в красный, белый или синий цвета, отсортируйте их по месту так,
чтобы объекты одного цвета были смежными, с цветами в порядке красного, белого и синего.
Мы будем использовать целые числа 0, 1 и 2 для представления красного, белого и синего цветов соответственно.
Вы должны решить эту проблему, не используя функцию сортировки библиотеки.

Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]


Выполнено:
Время выполнения: 0 мс, быстрее, чем 100% отправок в режиме онлайн.
Использование памяти: 2,1 МБ, менее 70.89% отправленных онлайн-заявок.
*/

package main

import "fmt"

func main() {
	//nums := []int{2, 0, 2, 1, 1, 0}
	nums := []int{2, 1, 0}
	fmt.Println(nums)

	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func sortColors1(nums []int) []int {
	n0 := make([]int, 0, len(nums))
	n1 := make([]int, 0, len(nums))
	n2 := make([]int, 0, len(nums))

	for _, d := range nums {
		switch d {
		case 0:
			n0 = append(n0, d)
		case 1:
			n1 = append(n1, d)
		case 2:
			n2 = append(n2, d)
		}
	}

	nums = make([]int, 0, len(nums))

	nums = append(nums, n0...)
	nums = append(nums, n1...)
	nums = append(nums, n2...)

	return nums
}
