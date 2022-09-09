/* https://leetcode.com/problems/add-binary/
67. Add Binary
Учитывая две двоичные строки a и b, верните их сумму в виде двоичной строки.

Пример 1:
Input: a = "11", b = "1"
Output: "100"

Пример 2:
Input: a = "1010", b = "1011"
Output: "10101"

Выполнено:
Время выполнения: 5 мс, быстрее, чем 25,27% онлайн-заявок на добавление двоичного файла.
Использование памяти: 3 МБ, менее 8,19% отправленных онлайн-заявок на добавление двоичного файла.

*/

package main

import "fmt"

func main() {
	//a, b := "1010", "1011"
	//a, b := "11", "1"
	//a, b := "1111", "1111"
	a, b := "1111", "00000"
	res := addBinary(a, b)
	fmt.Println("res =", res)

}

func addBinary(a string, b string) string {
	const base = 2
	aa := strInDigit(a)
	bb := strInDigit(b)
	n := len(aa)
	m := len(bb)

	i, j := 0, 0
	d := 0
	res := make([]int, 0, n+m)
	for i < n || j < m {
		if i < n && j < m {
			d += aa[i] + bb[j]
		}
		if i >= n {
			d += bb[j]
		}
		if j >= m {
			d += aa[i]
		}
		res = append(res, d%base)
		d /= base
		i++
		j++
	}
	if d > 0 {
		res = append(res, d)
	}

	return digitInStr(res)
}

func strInDigit(s string) []int {
	n := len(s)
	res := make([]int, 0, n)
	for i := n - 1; i > -1; i-- {
		if s[i] == '0' {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	}
	return res
}

func digitInStr(d []int) string {
	n := len(d)
	s := ""
	for i := n - 1; i > -1; i-- {
		if d[i] == 0 {
			s += "0"
		} else {
			s += "1"
		}
	}
	return s
}
