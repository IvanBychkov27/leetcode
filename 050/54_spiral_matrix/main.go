/* https://leetcode.com/problems/spiral-matrix/
54. Spiral Matrix
Учитывая матрицу m x n, верните все элементы матрицы в спиральном порядке.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

*/

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} //  [1,2,3,6,9,8,7,4,5]
	res := spiralOrder(matrix)
	fmt.Println("res:", res)
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix[0])
	n := len(matrix)
	res := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		res = append(res, matrix[0][i])
	}
	for i := 1; i < n; i++ {
		res = append(res, matrix[i][m-1])
	}
	for i := m - 2; i > -1; i-- {
		res = append(res, matrix[n-1][i])
	}
	for i := n - 2; i > 0; i-- {
		res = append(res, matrix[i][0])
	}

	for i := 0 + 1; i < m-1; i++ {
		res = append(res, matrix[1][i])
	}

	return res
}
