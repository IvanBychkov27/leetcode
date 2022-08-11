//https://leetcode.com/problems/sqrtx/
/*
Учитывая неотрицательное целое число x, вычислите и верните квадратный корень из x.
Поскольку возвращаемый тип является целым числом, десятичные цифры усекаются, и возвращается только целочисленная часть результата.
Примечание: Вам не разрешается использовать какие-либо встроенные функции или операторы экспоненты, такие как pow(x, 0.5) или x ** 0.5.

Выполнено:
Время выполнения: 26 мс, быстрее, чем 17,19% онлайн-заявок на Sqrt(x).
Использование памяти: 2,1 МБ, менее 23,44% заявок на участие в Sqrt(x).
*/
package main

import "fmt"

func main() {
	//x := 4
	x := 270

	res := mySqrt(x)
	fmt.Printf("sqrt(%d) = %d\n", x, res)
}

// sqrt(a^2+b) =(приближенно) a+b/2a
func mySqrt(x int) int {
	var a, d int
	if x >= 100 {
		a = 10
	}
	if x >= 10000 {
		a = 100
	}
	if x >= 1000000 {
		a = 1000
	}
	if x >= 100000000 {
		a = 10000
	}
	if x >= 10000000000 {
		a = 100000
	}
	if x >= 1000000000000 {
		a = 1000000
	}

	for {
		a++
		d = a * a
		if d == x {
			break
		}

		if d > x {
			a--
			break
		}
	}

	return a
}
