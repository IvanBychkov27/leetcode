/* https://leetcode.com/problems/max-points-on-a-line/
149. Max Points on a Line - Максимальное количество точек на линии
Учитывая массив точек, где точки [i] = [xi, yi] представляют точку на плоскости X-Y,
верните максимальное количество точек, которые лежат на одной прямой.

Пример 1:
Input: points = [[1,1],[2,2],[3,3]]
Output: 3

Пример 2:
Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4

Выполнено:
Время выполнения: 475 мс, быстрее, чем 6,10% заявок Go online для получения максимального количества очков в строке.
Использование памяти: 215,7 МБ, менее 6,10% заявок Go online на максимальное количество очков в строке.

*/

package main

import (
	"fmt"
)

func main() {
	//points := [][]int{{1, 1}, {2, 2}, {3, 3}} // 3
	//points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}} // 4
	//points := [][]int{{0, 0}} // 1
	//points := [][]int{{1, 0}, {0, 0}} // 2
	//points := [][]int{{0, 0}, {1, -1}, {1, 1}} // 2
	//points := [][]int{{0, 0}, {-1, -1}, {1, 1}} // 3
	//points := [][]int{{2, 3}, {3, 3}, {-5, 3}} // 3
	//points := [][]int{{2, -3}, {2, 0}, {2, 3}} // 3
	points := [][]int{{7, 3}, {19, 19}, {-16, 3}, {13, 17}, {-18, 1}, {-18, -17}, {13, -3}, {3, 7}, {-11, 12}, {7, 19}, {19, -12}, {20, -18}, {-16, -15}, {-10, -15}, {-16, -18}, {-14, -1}, {18, 10}, {-13, 8}, {7, -5}, {-4, -9}, {-11, 2}, {-9, -9}, {-5, -16}, {10, 14}, {-3, 4}, {1, -20}, {2, 16}, {0, 14}, {-14, 5}, {15, -11}, {3, 11}, {11, -10}, {-1, -7}, {16, 7}, {1, -11}, {-8, -3}, {1, -6}, {19, 7}, {3, 6}, {-1, -2}, {7, -3}, {-6, -8}, {7, 1}, {-15, 12}, {-17, 9}, {19, -9}, {1, 0}, {9, -10}, {6, 20}, {-12, -4}, {-16, -17}, {14, 3}, {0, -1}, {-18, 9}, {-15, 15}, {-3, -15}, {-5, 20}, {15, -14}, {9, -17}, {10, -14}, {-7, -11}, {14, 9}, {1, -1}, {15, 12}, {-5, -1}, {-17, -5}, {15, -2}, {-12, 11}, {19, -18}, {8, 7}, {-5, -3}, {-17, -1}, {-18, 13}, {15, -3}, {4, 18}, {-14, -15}, {15, 8}, {-18, -12}, {-15, 19}, {-9, 16}, {-9, 14}, {-12, -14}, {-2, -20}, {-3, -13}, {10, -7}, {-2, -10}, {9, 10}, {-1, 7}, {-17, -6}, {-15, 20}, {5, -17}, {6, -6}, {-11, -8}} // 6
	//points := [][]int{{5151, 5150}, {0, 0}, {5152, 5151}} // 2
	//points := [][]int{{5000000151, 5000000150}, {0, 0}, {5000000152, 5000000151}} // 2
	count := maxPoints(points)
	fmt.Println("max points on a line =", count)
}

func maxPoints(points [][]int) int {
	const (
		x = 0
		y = 1
	)

	if len(points) < 3 {
		return len(points)
	}

	arraysPoints := make(map[string][][]int)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			var k, b float64
			k = 1000000000
			if (points[j][x] - points[i][x]) != 0 {
				k *= float64((points[j][y] - points[i][y])) / float64((points[j][x] - points[i][x]))
				b = float64(points[i][y]) - k*float64(points[i][x])
			} else {
				k += float64(points[j][x])
			}
			key := fmt.Sprintf("%f@%f", k, b)
			d, ok := arraysPoints[key]
			if !ok {
				d = make([][]int, 0, len(points))
			}

			if !checkMap(d, points[i]) {
				d = append(d, points[i])
				arraysPoints[key] = d
			}
			if !checkMap(d, points[j]) {
				d = append(d, points[j])
				arraysPoints[key] = d
			}
			//fmt.Println(points[i], " ", points[j], " k =", k, " b =", b)
			//fmt.Printf("[%d %d] k = %0.64f b = %f\n", points[i], points[j], k, b)
		}
	}

	//fmt.Println(arraysPoints)
	return maxCountPoints(arraysPoints)
}

// нахождение максимального количества точек в массивах
func maxCountPoints(arraysPoints map[string][][]int) int {
	max := 2
	for _, points := range arraysPoints {
		l := len(points)
		if max < l {
			a := points[0]
			b := points[1]
			var ok bool
			for i := 2; i < l; i++ {
				ok = checkLine(a, b, points[i])
				if !ok {
					break
				}
			}
			if ok {
				max = l
			}
		}
	}
	return max
}

// проверка принадлежности трех точек одной прямой
func checkLine(a, b, c []int) bool {
	const (
		x = 0
		y = 1
	)
	if (a[x] == b[x] && b[x] == c[x]) || (a[y] == b[y] && b[y] == c[y]) {
		return true
	}

	if float64(c[x]-a[x])/float64(b[x]-a[x]) == float64(c[y]-a[y])/float64(b[y]-a[y]) {
		return true
	}
	return false
}

// true - если уже есть такой элемент в массиве
func checkMap(ds [][]int, b []int) bool {
	for _, d := range ds {
		if equalityArrays(d, b) {
			return true
		}
	}
	return false
}

// true - если массивы равны
func equalityArrays(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// --------------------------------
func maxPoints_01(points [][]int) int {
	const (
		x = 0
		y = 1
	)

	if len(points) < 3 {
		return len(points)
	}

	arraysPoints := make(map[int][][]int)
	for idx := 0; idx < len(points); idx++ {
		for i := 0; i < len(points); i++ {
			if idx == i {
				continue
			}
			for j := 0; j < len(points); j++ {
				if j == i || j == idx {
					continue
				}

				k := float64(100000)
				if (points[j][x] - points[i][x]) != 0 {
					k = float64((points[i][y] - points[idx][y])) / float64((points[i][x] - points[idx][x]))
				}

				key := idx + int(1000000*k)
				d, ok := arraysPoints[key]
				if !ok {
					d = make([][]int, 0, len(points))
				}
				if checkLine(points[idx], points[i], points[j]) {
					arraysPoints[key] = addArray(d, points[idx], points[i], points[j])
				}
			}
		}
	}

	//fmt.Println(arraysPoints)
	count := maxCount(arraysPoints)
	if count == 0 {
		count = 2
	}
	return count
}

// нахождение максимального количества точек в массивах
func maxCount(arraysPoints map[int][][]int) int {
	var max int
	for _, points := range arraysPoints {
		l := len(points)
		if max < l {
			max = l
		}
	}
	return max
}

// функция добавление только новых элементов в массив
func addArray(ds [][]int, b ...[]int) [][]int {
	for _, d := range b {
		if !checkMap(ds, d) {
			ds = append(ds, d)
		}
	}

	return ds
}

//func maxPoints(points [][]int) int {
//	const (
//		x = 0
//		y = 1
//	)
//
//	var count int
//	dK := make(map[float64][][]int)
//	for i := 0; i < len(points)-1; i++ {
//		for j := i + 1; j < len(points); j++ {
//			k := float64(10000)
//			if (points[j][x] - points[i][x]) != 0 {
//				k = float64((points[j][y] - points[i][y])) / float64((points[j][x] - points[i][x]))
//			} else {
//				k += float64(points[j][x])
//			}
//
//			if (points[j][y] - points[i][y]) == 0 {
//				k = float64(1000 + points[j][y])
//			}
//
//			d, ok := dK[k]
//			if !ok {
//				d = make([][]int, 0, len(points))
//			}
//
//			if !checkMap(d, points[i]) {
//				d = append(d, points[i])
//				dK[k] = d
//			}
//			if !checkMap(d, points[j]) {
//				d = append(d, points[j])
//				dK[k] = d
//			}
//
//			fmt.Println(points[i], " ", points[j], " k =", k)
//		}
//	}
//
//	fmt.Println(dK)
//
//	return count
//}
