// https://leetcode.com/problems/reverse-integer/
/*
Учитывая 32-разрядное целое число x со знаком, верните x с обратными цифрами. Если обращение x приводит к выходу значения за пределы 32-разрядного целого диапазона со знаком [-231, 231 - 1], то верните 0.
Предположим, что среда не позволяет хранить 64-разрядные целые числа (со знаком или без знака).

Результат:
Время выполнения: 4 мс, быстрее, чем 37,23 % онлайн-отправлений для обратного целого числа.
Использование памяти: 2,3 МБ, менее 16,33 % онлайн-заявок на обратное целое число.


Результат 3:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправлений в режиме онлайн для обратного целого числа.
Использование памяти: 2,2 МБ, менее 100,00 % отправленных онлайн-сообщений для обратного целого числа.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := -214748399
	//x := 1020100
	//x := -120
	res := reverse(x)
	fmt.Println("res  =", res)

	res2 := reverse2(x)
	fmt.Println("res2 =", res2)

	res3 := reverse3(x)
	fmt.Println("res3 =", res3)

}

func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}

	min := 1
	if x < 0 {
		min = -1
		x *= min
	}

	s := strconv.Itoa(x)
	r := make([]string, len(s))
	for i, v := range s {
		r[len(s)-1-i] = string(v)
	}

	d := ""
	for _, v := range r {
		d += v
	}

	var res int
	res, _ = strconv.Atoi(d)

	res *= min

	if res < -2147483648 || res > 2147483647 {
		res = 0
	}

	return res
}

func reverse2(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}

	min := 1
	if x < 0 {
		min = -1
		x *= min
	}

	data := make([]int, 0, 10)
	for {
		d := x % 10
		data = append(data, d)
		x /= 10
		if x == 0 {
			break
		}
	}

	var res int
	l := len(data)
	m := 10
	j := l
	for i := 0; i < l; i++ {
		j--
		d := data[i]
		if d == 0 {
			continue
		}
		res += int(math.Pow(float64(m), float64(j))) * d
	}

	res *= min

	if res < -2147483648 || res > 2147483647 {
		res = 0
	}

	return res
}

func reverse3(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}

	min := 1
	if x < 0 {
		min = -1
		x *= min
	}

	l := 0
	temp := x
	for {
		l++
		temp /= 10
		if temp == 0 {
			break
		}
	}

	var res int
	for {
		l--
		d := x % 10
		if d != 0 {
			res += int(math.Pow(float64(10), float64(l))) * d
		}
		x /= 10
		if x == 0 {
			break
		}
	}

	res *= min

	if res < -2147483648 || res > 2147483647 {
		res = 0
	}

	return res
}
