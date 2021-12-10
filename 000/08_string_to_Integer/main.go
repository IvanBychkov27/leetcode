// https://leetcode.com/problems/string-to-integer-atoi/
/*
 Реализуйте функцию myAtoi (string s), которая преобразует строку в 32-разрядное целое число со знаком (аналогично функции atoi C/C++).

Результат:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправлений в режиме онлайн для преобразования строки в целое число (atoi).
Использование памяти: 2,3 МБ, менее 100,00 % онлайн-отправлений для преобразования строки в целое число (atoi).
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	s := " -4193 with words" // -4193
	//s := "42"
	//s := "words and 987"
	//s := "-91283472332"
	//s := "+-12"
	//s := "  0000000000012345678"
	//s := "   +0 123"
	//s := "2000000000000000000000"
	//s := "  +  413"

	res := myAtoi(s)
	fmt.Println("res  =", res)

	res2 := myAtoi(s)
	fmt.Println("res2 =", res2)

}

func myAtoi(s string) int {
	data := []byte(s)
	zn := 1
	f := false

	t := make([]int, 0, 64)
	for _, d := range data {
		if d == ' ' && len(t) == 0 {
			if f {
				return 0
			}
			continue
		}
		if d == '-' && len(t) == 0 {
			if f {
				return 0
			}
			zn = -1
			f = true
			continue
		}
		if d == '+' && len(t) == 0 {
			if f {
				return 0
			}
			zn = 1
			f = true
			continue
		}
		if (d < 48 || d > 57) && len(t) == 0 {
			return 0
		}

		if d >= 48 && d <= 57 {
			t = append(t, int(d-48))
		} else {
			break
		}
	}

	var res int
	l := len(t)
	m := 10
	j := l
	for i := 0; i < l; i++ {
		j--
		d := t[i]
		if d == 0 {
			continue
		}

		if j > 11 {
			if zn < 0 {
				return -2147483648
			} else {
				return 2147483647
			}
		}

		res += int(math.Pow(float64(m), float64(j))) * d
	}

	res *= zn

	if res < -2147483648 {
		res = -2147483648
	}
	if res > 2147483647 {
		res = 2147483647
	}

	return res
}

func myAtoi_2(s string) int {
	const (
		min = -2147483648
		max = 2147483647
	)
	data := []byte(s)
	zn := 1
	f := false

	t := make([]int, 0, 64)
	for _, d := range data {
		if len(t) == 0 && (d == (' ') || d == '-' || d == '+') {
			if f {
				return 0
			}
			switch d {
			case ' ':
			case '-':
				zn = -1
				f = true
			case '+':
				zn = 1
				f = true
			}
			continue
		}

		if (d < 48 || d > 57) && len(t) == 0 {
			return 0
		}

		if d >= 48 && d <= 57 {
			t = append(t, int(d-48))
		} else {
			break
		}
	}

	var res int
	l := len(t)
	m := 10
	j := l
	for i := 0; i < l; i++ {
		j--
		d := t[i]
		if d == 0 {
			continue
		}

		if j > 11 {
			if zn < 0 {
				return min
			} else {
				return max
			}
		}

		res += int(math.Pow(float64(m), float64(j))) * d
	}

	res *= zn

	if res < min {
		res = min
	}
	if res > max {
		res = max
	}

	return res
}
