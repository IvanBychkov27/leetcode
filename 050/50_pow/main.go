/* https://leetcode.com/problems/powx-n/

Реализуйте pow(x, n), который вычисляет x, возведенный в степень n

Выполнено:
Время выполнения: 2599 мс, быстрее, чем 6,30% онлайн-заявок для Pow(x, n).
Использование памяти: 2 МБ, менее 19,45% заявок на участие в конкурсе Pow(x, n).
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2.0
	n := 7
	fmt.Printf("%.0f ^ %d = %f\n", x, n, math.Pow(x, float64(n)))
	fmt.Printf("%.0f ^ %d = %f\n", x, n, myPow(x, n))
}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 || x == 1.0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	if n < -100000 {
		return 0
	}

	k := n
	if n < 0 {
		k = -n
	}

	res := x
	for i := 1; i < k; i++ {
		res *= x
	}

	if n < 0 {
		res = 1 / res
	}

	return res
}
