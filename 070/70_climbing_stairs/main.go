/* https://leetcode.com/problems/climbing-stairs/
70. Climbing Stairs - Подъем По Лестнице

Вы поднимаетесь по лестнице. Чтобы достичь вершины, требуется n шагов.
Каждый раз вы можете подняться либо на 1, либо на 2 ступеньки. Сколькими различными способами вы можете подняться на вершину?

Решение сводится к подсчету числа Фибоначчи на шаге n,
только в данной задаче вначале необходимо n увеличить на 1 (n = n +1)

Пример 1:
	Input: n = 2
	Output: 2
	Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps

Пример 2:
	Input: n = 3
	Output: 3
	Explanation: There are three ways to climb to the top.
	1. 1 step + 1 step + 1 step
	2. 1 step + 2 steps
	3. 2 steps + 1 step

Выполнено:
Время выполнения: 1 мс, быстрее, чем 54,72% онлайн-заявок на подъем по лестнице.
Использование памяти: 1,9 МБ, менее 91,29% заявок Go online на подъем по лестнице.
*/
package main

import "fmt"

func main() {
	//n := 2
	n := 3
	res := climbStairs(n)
	fmt.Println("res = ", res)
}

func climbStairs(n int) int {
	var a, b, c int
	a = 1 // 1е число Фибоначчи = 1
	b = 1 // 2е число Фибоначчи = 1
	for i := 1; i < n; i++ {
		c = a + b // последующее число Фибоначчи
		a = b
		b = c
	}
	return c
}

// Числа Фибоначчи
// 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229 832040 1346269 2178309 3524578 5702887 9227465 14930352 24157817 39088169 63245986 102334155
