/* https://leetcode.com/problems/merge-intervals/
56. Merge Intervals
Учитывая массив интервалов, где интервалы [i] = [starti, endi], объедините все перекрывающиеся интервалы и
верните массив неперекрывающихся интервалов, которые охватывают все интервалы во входных данных.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Пояснение: Поскольку интервалы [1,3] и [2,6] перекрываются, объедините их в [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Пояснение: Интервалы [1,4] и [4,5] считаются перекрывающимися.

Выполнено:
Время выполнения: 468 мс, быстрее, чем 5,14% отправок в режиме онлайн для интервалов слияния.
Использование памяти: 6,7 МБ, менее 90,66% отправок в режиме онлайн для интервалов слияния.

*/

package main

import "fmt"

func main() {
	//intervals := [][]int{{1, 4}, {4, 5}}
	//intervals := [][]int{{5, 9}, {4, 7}}
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{0, 1}, {2, 6}, {5, 10}, {9, 18}}
	//intervals := [][]int{{1, 4}, {0, 0}}
	res := merge(intervals)
	fmt.Println("res:", res)

	//fmt.Println(isIntersection([]int{1, 4}, []int{4, 7}))
}

func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			tmp, ok := isIntersection(intervals[i], intervals[j])
			if ok {
				intervals[i] = nil
				intervals[j] = tmp
			}
		}
	}

	res := make([][]int, 0, len(intervals))
	for _, d := range intervals {
		if d != nil {
			res = append(res, d)
		}
	}
	return res
}

func isIntersection(a, b []int) ([]int, bool) {
	if len(a) != 2 || len(b) != 2 {
		return nil, false
	}
	var i, j int
	if b[0] < a[0] {
		a, b = b, a
	}
	if b[0] <= a[1] {
		i = a[0]
		j = b[1]
		if b[1] < a[1] {
			j = a[1]
		}
		return []int{i, j}, true
	}
	return nil, false
}
