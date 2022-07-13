/* https://leetcode.com/problems/divide-two-integers/

Учитывая два целых числа, делимое и делитель, разделите два целых числа без использования оператора умножения, деления и модуляции.
Целочисленное деление должно быть усечено до нуля, что означает потерю его дробной части. Например, 8.345 будет усечено до 8, а -2.7335 будет усечено до -2.

Верните частное после деления дивиденда на делитель.
Примечание: Предположим, что мы имеем дело со средой, которая может хранить только целые числа в пределах 32-разрядного диапазона целых чисел со знаком: [-2^31, 2^31 − 1]. 2147483648
Для этой задачи, если частное строго больше 2^31 - 1, то верните 231 - 1, а если частное строго меньше -2^31, то верните -2^31.

Выполнено:
Время выполнения: 15 мс, быстрее, чем 28,73% онлайн-заявок на Раздел двух целых чисел.
Использование памяти: 2,3 МБ, менее 89,26% заявок Go online на Раздел двух целых чисел.

*/
package main

import "fmt"

func main() {
	dividend := -2147483648
	divisor := -3

	res := divide(dividend, divisor)
	fmt.Println("RES =", dividend/divisor)
	fmt.Println("res =", res)
}

func divide(dividend int, divisor int) int {
	var z bool
	if dividend < 0 {
		dividend = -dividend
		z = !z
	}
	if divisor < 0 {
		divisor = -divisor
		z = !z
	}

	var res int
	div := divisor
	const n = 10000
	if dividend > n {
		div = divisor * n
		for {
			if dividend < div {
				break
			}
			res++
			dividend -= div
		}
		res *= n
	}

	for {
		if dividend < divisor {
			break
		}
		res++
		dividend -= divisor
	}

	if z {
		res = -res
	}

	if res >= 2147483648 {
		res = 2147483648 - 1
	}

	if res < -2147483648 {
		res = -2147483648
	}

	return res
}
