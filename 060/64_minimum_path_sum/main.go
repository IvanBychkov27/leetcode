/* https://leetcode.com/problems/minimum-path-sum/
64. Minimum Path Sum
Учитывая сетку m x n, заполненную неотрицательными числами, найдите путь от верхнего левого угла до нижнего правого,
который минимизирует сумму всех чисел на своем пути.
Примечание: Вы можете перемещаться только вниз или вправо в любой момент времени.

Example 1:
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Example 2:
Input: grid = [[1,2,3],[4,5,6]]
Output: 12

Выполнено:
Время выполнения: 20 мс, быстрее, чем 21,96% отправок в режиме онлайн при минимальной сумме пути.
Использование памяти: 4,4 МБ, менее 48,24% отправленных онлайн-заявок на минимальную сумму пути.

*/
package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}} // 7
	//grid := [][]int{{1, 2, 3}, {4, 5, 6}} // 12
	//grid := [][]int{{1, 2}, {4, 5}, {1, 2}, {4, 6}} // 14

	printField(grid)

	res := minPathSum(grid)
	fmt.Println("sum min path:", res)
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	field := make([][]int, n)
	d := make([]int, m)
	d[0] = grid[0][0]
	for j := 1; j < m; j++ {
		d[j] = grid[0][j] + d[j-1]
	}
	field[0] = d

	for i := 1; i < n; i++ {
		d = make([]int, m)
		d[0] = grid[i][0] + field[i-1][0]
		field[i] = d
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			field[i][j] = grid[i][j] + min(field[i][j-1], field[i-1][j])
		}
	}

	return field[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printField(field [][]int) {
	for _, d := range field {
		fmt.Println(d)
	}
}
