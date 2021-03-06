/* https://leetcode.com/problems/roman-to-integer/
От римского к целому числу

Римские цифры представлены семью различными символами:  I, V, X, L, C, D, M.
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

Например, 2 записывается как II римской цифрой, просто две единицы складываются вместе. 12 записывается как XII, то есть просто X + II.
Число 27 записывается как XXVII, то есть XX + V + II.

Римские цифры обычно пишутся слева направо от наибольшего к наименьшему. Однако цифра для четырех не является IIII.
Вместо этого число четыре записывается как IV. Поскольку единица стоит перед пятью, мы вычитаем ее, получая четыре.
Тот же принцип применим и к числу девять, которое записывается как IX. Существует шесть случаев, когда используется вычитание:

I м.б. помещен перед V (5) и X (10), чтобы сделать 4 и 9.
X м.б. помещен перед L (50) и C (100), чтобы сделать 40 и 90.
C м.б. помещен перед D (500) и M (1000), чтобы сделать 400 и 900.

Учитывая римскую цифру, преобразуйте ее в целое число.

Выполнено:
Время выполнения: 4 мс, быстрее, чем 89,73 % онлайн-отправлений для перевода с римской цифры на целочисленный.
Использование памяти: 3 МБ, менее 45,82 % онлайн-заявок на перевод с римской цифры на целочисленный.

*/
package main

import (
	"fmt"
)

func main() {
	s := "MCMXCIV" // 1994

	res := romanToInt(s)
	fmt.Println(s, "=", res)
}

func romanToInt(s string) int {
	if len(s) == 0 || len(s) > 15 {
		return 0
	}
	var p byte
	res := 0
	data := []byte(s)
	for _, d := range data {
		switch d {
		case 'I':
			res += 1
		case 'V':
			if p == 'I' {
				res += 3
				break
			}
			res += 5
		case 'X':
			if p == 'I' {
				res += 8
				break
			}
			res += 10
		case 'L':
			if p == 'X' {
				res += 30
				break
			}
			res += 50
		case 'C':
			if p == 'X' {
				res += 80
				break
			}
			res += 100
		case 'D':
			if p == 'C' {
				res += 300
				break
			}
			res += 500
		case 'M':
			if p == 'C' {
				res += 800
				break
			}
			res += 1000
		}
		p = d
	}

	return res
}
