// https://leetcode.com/problems/search-a-2d-matrix/
/* 74. Search a 2D Matrix
Напишите эффективный алгоритм, который ищет целевое значение в целочисленной матрице m x n matrix. Эта матрица обладает следующими свойствами:
Целые числа в каждой строке сортируются слева направо. Первое целое число каждой строки больше последнего целого числа предыдущей строки.

Example 1:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

Выполнено:
Время выполнения: 4 мс, быстрее, чем 69.48% отправок в режиме онлайн.
Использование памяти: 2,7 МБ, менее 75.55% отправленных онлайн-заявок.

*/

package main

import "fmt"

func main() {
	//matrix, target := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3
	//matrix, target := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13
	//matrix, target := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 65, 67, 70}}, 21
	matrix, target := [][]int{{0}}, 1

	res := searchMatrix(matrix, target)
	fmt.Println(target, res)
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	if m == 0 || n == 0 {
		return false
	}

	var i, j, x, y int
	y = m
	for {
		j = (x + y) / 2
		d := matrix[j][0]
		if d == target {
			return true
		}
		if target < d {
			y = j
		} else {
			x = j
		}
		if (x+1) == y || x == y {
			j = x
			break
		}
	}

	x, y = 0, n
	for {
		i = (x + y) / 2
		d := matrix[j][i]
		if d == target {
			return true
		}
		if target < d {
			y = i
		} else {
			x = i
		}
		if (x+1) == y || x == y {
			i = x
			break
		}
	}

	if matrix[j][i] == target {
		return true
	}

	return false
}
