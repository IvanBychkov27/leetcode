/* https://leetcode.com/problems/valid-number/
65. Valid Number (Hard)
Действительное число может быть разделено на эти компоненты (по порядку):
Десятичное число или целое число.
(опционально) Буква "е" или "E", за которой следует целое число.

Десятичное число может быть разделено на эти компоненты (по порядку):
(опционально) Знаковый символ (либо '+', либо '-').
Один из следующих форматов:
Одна или несколько цифр, за которыми следует точка ".".
Одна или несколько цифр, за которыми следует точка ".", за которой следует одна или несколько цифр.
Точка ".", за которой следует одна или несколько цифр.

Целое число может быть разделено на эти компоненты (по порядку):
(опционально) Знаковый символ (либо '+', либо '-').
Одна или несколько цифр.

Например, все приведенные ниже действительные числа: ["2", "0089", "-0.1", "+3.14", "4.", "- .9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"],
в то время как следующие недопустимые номера: ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", " 95a54e53"].

Учитывая строку s, верните значение true, если s является допустимым числом.

Example 1: Input: s = "0" Output: true
Example 2: Input: s = "e" Output: false
Example 3: Input: s = "." Output: false

Ограничения:
1 <= s.length <= 20
s состоит только из английских букв (как прописных, так и строчных), цифр (0-9), плюс '+', минус '-' или точка '.'.

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100% отправок в режиме онлайн.
Использование памяти: 2,4 МБ, менее 14,75% отправленных онлайн-заявок.

*/

package main

import "fmt"

func main() {
	s := "46.e3" // true
	res := isNumber(s)
	fmt.Printf("%s -> %t", s, res)
}

func isNumber(s string) bool {
	if !verification(s) {
		return false
	}

	return true
}

func verification(s string) bool {
	control := map[string]bool{
		"e": true,
		"E": true,
		".": true,
		"-": true,
		"+": true,
	}

	pred := ""
	mantissa := ""
	for _, ds := range s {
		d := string(ds)
		switch d {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			mantissa += d
			control["-"] = false
			control["+"] = false

		case "e", "E":
			if !control[d] || pred == "" || mantissa == "" {
				return false
			}
			control["e"] = false
			control["E"] = false
			control["."] = false
			control["-"] = true
			control["+"] = true

		case ".":
			if !control[d] || pred == "e" || pred == "E" {
				return false
			}
			control[d] = false

		case "-":
			if !control[d] || pred == "+" || pred == "." {
				return false
			}
			control[d] = false

		case "+":
			if !control[d] || pred == "-" || pred == "." {
				return false
			}
			control[d] = false
		default:
			return false
		}
		pred = d
	}

	if pred == "e" || pred == "E" || pred == "-" || pred == "+" || mantissa == "" {
		return false
	}

	return true
}
