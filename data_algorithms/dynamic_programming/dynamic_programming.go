/* https://www.youtube.com/watch?v=4E-_uzO0A_A&list=PLRDzFCPr95fL_5Xvnufpwj2uYZnZBBnsr&index=13
	Двумерное динамическое программирование
Задача подсчета кол-ва возможных путей муровья из верхней левой клетки таблицы в ячейку с координатами [n,m]
При условии когда муравей может перемещаться на каждом шаге только на одну клетку вправо или влево

*/

package main

import (
	"fmt"
)

func main() {
	n, m := 5, 4
	res := dynamicProcess(n, m)
	fmt.Println("res =", res)
}

func dynamicProcess(n, m int) int {
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
