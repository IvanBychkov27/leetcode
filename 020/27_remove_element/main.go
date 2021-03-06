/* https://leetcode.com/problems/remove-element/

Учитывая целочисленный массив nums и целое значение val, удалите все вхождения val в nums на месте. Относительный порядок элементов может быть изменен.
Поскольку в некоторых языках невозможно изменить длину массива, вместо этого вы должны поместить результат в первую часть массива nums.
Более формально, если после удаления дубликатов осталось k элементов, то первые k элементов nums должны содержать конечный результат.
Не имеет значения, что вы оставляете за пределами первых k элементов.

Верните k после размещения конечного результата в первых k слотах nums.
Не выделяйте дополнительное пространство для другого массива. Вы должны сделать это, изменив входной массив на месте с помощью O(1) дополнительной памяти.

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Объяснение: Ваша функция должна возвращать k = 2, причем первые два элемента nums равны 2.
Не имеет значения, что вы оставляете за пределами возвращаемого k (следовательно, они являются символами подчеркивания).

Выполнено:
Время выполнения: 2 мс, быстрее, чем 50,65% онлайн-заявок на удаление элемента.
Использование памяти: 2 МБ, менее 83,03% отправленных онлайн-заявок на удаление элемента.
*/
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	fmt.Println(nums)
	k := removeElement(nums, val)

	fmt.Println("del =", val)
	fmt.Print(nums[:k])
	fmt.Println("  len =", k)
}

func removeElement(nums []int, val int) int {
	var idx int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
