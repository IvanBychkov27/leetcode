// https://leetcode.com/problems/zigzag-conversion/
/*
Строка "PAYPALISHIRING" написана зигзагообразным узором на заданном количестве строк следующим образом: (возможно, вы захотите отобразить этот узор фиксированным шрифтом для лучшей разборчивости)
А затем прочитайте строчку за строчкой: "PAHNAPLSIIGYIR".
Напишите код, который будет принимать строку, и выполните это преобразование с учетом количества строк

Результат:
Время выполнения: 12 мс, быстрее, чем 57,38 % отправленных онлайн-сообщений для преобразования зигзага.
Использование памяти: 7,3 МБ, что составляет менее 44,96 % от отправленных онлайн сообщений для преобразования зигзага.
*/
package main

import (
	"fmt"
)

func main() {
	//s := "PAYPALISHIRING"
	//numRows := 3
	//exected := "PAHNAPLSIIGYIR"

	//s := "PAYPALISHIRING"
	//numRows := 4
	//exected := "PINALSIGYAHRPI"

	s := "AB"
	numRows := 1
	exected := "AB"

	res := convert(s, numRows)
	fmt.Println("res =", res)
	fmt.Println(res == exected)

}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var f bool
	data := make([]string, numRows)
	ss := []byte(s)
	idx := 0
	for _, v := range ss {
		data[idx] += string(v)

		if !f {
			idx++
		} else {
			idx--
		}

		if idx == 0 {
			f = false
		}
		if idx == numRows-1 {
			f = true
		}

	}

	res := ""
	for _, v := range data {
		res += v
	}

	return res
}
