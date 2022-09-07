/*  https://leetcode.com/problems/multiply-strings/
43. Multiply Strings - Умножение строк
	Учитывая два неотрицательных целых числа num1 и num2, представленных в виде строк, верните произведение num1 и num2, также представленное в виде строки.
	Примечание: Вы не должны использовать какую-либо встроенную библиотеку BigInteger или напрямую преобразовывать входные данные в целое число.

Пример 1:
Input: num1 = "2", num2 = "3"
Output: "6"

Пример 2:
Input: num1 = "123", num2 = "456"
Output: "56088"

Ограничения:
1 <= num1.length, num2.length <= 200
num1 и num2 состоят только из цифр.
Как num1, так и num2 не содержат никакого начального нуля, кроме самого числа 0.

Выполнено:
Время выполнения: 17 мс, быстрее, чем 24,85% онлайн-заявок для умножения строк.
Использование памяти: 3,2 МБ, менее 39,35% отправленных онлайн-заявок на умножение строк.
*/

package main

import "fmt"

func main() {
	num1, num2 := "0", "0"
	//num1, num2 := "2", "3"
	//num1, num2 := "123", "456"
	res := multiply(num1, num2)
	fmt.Println(res)
}

func multiply(num1 string, num2 string) string {
	const base = 10

	a := parseDataInArray(num1)
	b := parseDataInArray(num2)

	n := len(a)
	m := len(b)

	res := make([]int, n+m)
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			d := a[i] * b[j]
			res[i+j] += d % base
			res[i+j+1] += d / base
			if res[i+j] >= base {
				res[i+j], res[i+j+1] = res[i+j]%base, res[i+j+1]+(res[i+j]/base)
			}
		}
	}
	fmt.Println("res =", res)
	return printRes(res)
}

func parseDataInArray(num string) []int {
	const correct = 48
	var digit int
	n := len(num)
	res := make([]int, 0, n)
	for i := n - 1; i > -1; i-- {
		digit = int(num[i] - correct)
		if i == 0 && digit == 0 {
			continue
		}
		res = append(res, digit)
	}
	return res
}

func printRes(d []int) string {
	var s string
	f := true
	for i := len(d) - 1; i > -1; i-- {
		if f {
			if d[i] == 0 {
				continue
			}
			f = false
		}
		s += fmt.Sprintf("%d", d[i])
	}
	if s == "" {
		s = "0"
	}
	return s
}
