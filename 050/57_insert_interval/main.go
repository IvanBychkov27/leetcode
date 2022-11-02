/* https://leetcode.com/problems/insert-interval/
57. Insert Interval

Вам предоставляется массив неперекрывающихся интервалов intervals, где интервалы[i] = [starti, endi] представляют начало и конец i-го интервала,
а интервалы сортируются в порядке возрастания по starti.
Вам также предоставляется интервал newInterval = [start, end], который представляет начало и конец другого интервала.
Вставьте новый интервал в интервалы таким образом, чтобы интервалы по-прежнему сортировались в порядке возрастания по началу,
а интервалы по-прежнему не имели перекрывающихся интервалов (при необходимости объедините перекрывающиеся интервалы).
Возвращайте интервалы после вставки.

Example 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Объяснение: Поскольку новый интервал [4,8] перекрывается с [3,5],[6,7],[8,10].

Выполнено:
Время выполнения: 111 мс, быстрее, чем 5,52% онлайн-отправлений для интервала вставки.
Использование памяти: 5 МБ, менее 10,57% отправленных онлайн-заявок для интервала вставки.

*/

package main

import "fmt"

func main() {
	//intervals, newInterval := [][]int{{1, 3}, {6, 9}}, []int{2, 5} // [[1,5],[6,9]]
	//intervals, newInterval := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8} // [[1,2],[3,10],[12,16]]
	//intervals, newInterval := [][]int{}, []int{5, 7} // [[5 7]]
	//intervals, newInterval := [][]int{{1, 2}}, []int{5, 7} // [[1 2] [5 7]]
	//intervals, newInterval := [][]int{{8, 9}}, []int{5, 7} // [[5 7] [8 9]]
	//intervals, newInterval := [][]int{{1, 5}}, []int{2, 3} // [[1 5]]
	//intervals, newInterval := [][]int{{2, 3}}, []int{1, 7} // [[1 7]]
	intervals, newInterval := [][]int{{3, 5}, {12, 15}}, []int{6, 8} // [[3 5] [6 8] [12 15]]
	res := insert(intervals, newInterval)
	fmt.Println("res:", res)

	//fmt.Println(isIntersection(newInterval, intervals[0]))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	if newInterval[1] < intervals[0][0] {
		tmp := intervals
		intervals = [][]int{newInterval}
		intervals = append(intervals, tmp...)
		return intervals
	}

	if intervals[len(intervals)-1][1] < newInterval[0] {
		intervals = append(intervals, newInterval)
		return intervals
	}

	f := true
	for i := 0; i < len(intervals); i++ {
		if f {
			newInt, ok := isIntersection(newInterval, intervals[i])
			if ok {
				intervals[i] = newInt
				f = false
			}
		}

		if !f {
			for j := i + 1; j < len(intervals); j++ {
				tmp, ok := isIntersection(intervals[i], intervals[j])
				if ok {
					intervals[i] = nil
					intervals[j] = tmp
				}
			}
		}
	}

	res := make([][]int, 0, len(intervals))
	p := 0
	for _, d := range intervals {
		if f {
			if p < newInterval[0] && d[0] > newInterval[1] {
				res = append(res, newInterval)
				f = false
			}
			p = d[1]
		}

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
