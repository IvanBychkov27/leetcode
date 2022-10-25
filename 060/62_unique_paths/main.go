/* https://leetcode.com/problems/unique-paths/
62. Unique Paths - Уникальные пути
На сетке mxn есть робот. Робот изначально расположен в верхнем левом углу (т.е. сетка[0][0]). Робот пытается переместиться в правый нижний угол (т.е. в сетку [m - 1][n - 1]). Робот может двигаться только вниз или вправо в любой момент времени.
Учитывая два целых числа m и n, верните количество возможных уникальных путей, которые робот может пройти, чтобы добраться до нижнего правого угла.
Тестовые примеры генерируются таким образом, чтобы ответ был меньше или равен 2 * 109.

Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00% отправок в режиме онлайн для уникальных путей.
Использование памяти: 2,1 МБ, менее 58,67% отправленных онлайн-заявок на уникальные пути.

*/
package main

import "fmt"

func main() {
	m, n := 3, 7
	//m, n := 3, 2
	res := uniquePaths(m, n)
	fmt.Println("res =", res)
}

func uniquePaths(m int, n int) int {
	field := setField(n, m)
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			field[i][j] = field[i][j-1] + field[i-1][j]
		}
	}
	printField(field)
	return field[n-1][m-1]
}

func setField(n, m int) [][]int {
	field := make([][]int, n)

	d := make([]int, m)
	for j := 0; j < m; j++ {
		d[j] = 1
	}
	field[0] = d

	for i := 1; i < n; i++ {
		d = make([]int, m)
		d[0] = 1
		field[i] = d
	}
	return field
}

func printField(field [][]int) {
	for _, d := range field {
		fmt.Println(d)
	}
}
