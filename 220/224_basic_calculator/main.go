/* https://leetcode.com/problems/basic-calculator/
224. Basic Calculator - Базовый калькулятор
Учитывая строку s, представляющую допустимое выражение, реализуйте базовый калькулятор для ее вычисления и верните результат вычисления.
Примечание: Вам не разрешается использовать какие-либо встроенные функции, которые вычисляют строки как математические выражения, такие как eval().

Example 1:
Input: s = "1 + 1"
Output: 2

Example 2:
Input: s = " 2-1 + 2 "
Output: 3

Example 3:
Input: s = "(1+(4+5+2)-3)+(6+8)"
Output: 23

Ограничения:
1 <= s.длина <= 3 * 105
s состоит из цифр, '+', '-', '(', ')', и ' '.
s представляет собой допустимое выражение.
'+' не используется как унарная операция (т.е. "+1" и "+(2 + 3)" недопустимы).
'-' может использоваться как унарная операция (т.е. допустимы "-1" и "-(2 + 3)").
Во входных данных не будет двух последовательных операторов.
Каждое число и выполняемое вычисление будут помещаться в 32-разрядное целое число со знаком.

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//s := "-5+6-(7+2-(1+1+(3-1)+3-1))" //-2
	//s := "1+(2+1-(5-4))" // 3
	//s := "5 - (2 + 1)-(1-2)" // 3
	//s := "(1+(4+5+2)-3)+(6+8)" // 23
	//s := "20 - (1 + 2 + 22) - 18" // -23
	//s := " 2-1 + 2 " // 3
	//s := "10 - 3" // 7
	//s := "1 + 1" // 2
	//s := "0123456789"
	//s := "5+3-4-(1+2-7+(10-1+3+5+(3-0+(8-(3+(8-(10-(6-10-8-7+(0+0+7)-10+5-3-2+(9+0+(7+(2-(2-(9)-2+5+4+2+(2+9+1+5+5-8-9-2-9+1+0)-(5-(9)-(0-(7+9)+(10+(6-4+6))+0-2+(10+7+(8+(7-(8-(3)+(2)+(10-6+10-(2)-7-(2)+(3+(8))+(1-3-8)+6-(4+1)+(6))+6-(1)-(10+(4)+(8)+(5+(0))+(3-(6))-(9)-(4)+(2))))))-1)))+(9+6)+(0))))+3-(1))+(7))))))))" // -35
	s := "-(-10)-(-1+2)" // 10

	res := calculate(s)
	fmt.Printf("%s = %d\n", s, res)
}

func calculate(s string) int {
	steckOp = []int{}
	return sum(operations(condition(s)))
}

const (
	unknown    = 0
	digit      = 1
	operation  = 2
	bracket    = 3
	whitespace = 4
)

var steckOp []int

func sum(ds []int) int {
	fmt.Println(ds)
	res := 0
	for _, d := range ds {
		res += d
	}
	return res
}

func operations(cond []string) []int {
	ops := make([]int, 0, len(cond))
	op := 1
	for _, el := range cond {
		switch el {
		case "+":
			op = 1
		case "-":
			op = -1
		case "(":
			pushSteckOp(op * readSteckOp())
			op = 1
		case ")":
			getSteckOp()
		default:
			d := op * readSteckOp() * convertInDigit(el)
			ops = append(ops, d)
		}
	}
	return ops
}

func pushSteckOp(d int) {
	steckOp = append(steckOp, d)
}

func readSteckOp() int {
	if len(steckOp) == 0 {
		return 1
	}
	return steckOp[len(steckOp)-1]
}

func getSteckOp() int {
	if len(steckOp) == 0 {
		return 1
	}
	d := steckOp[len(steckOp)-1]
	steckOp = steckOp[:len(steckOp)-1]
	return d
}

// разбиение строки по элементам условия
func condition(s string) []string {
	res := make([]string, 0, len(s))
	s = clearWhitespace(s)
	d := ""
	f := digit
	for _, item := range s {
		el := element(item)
		if el != f {
			f = el
			res = append(res, d)
			d = ""
		}
		d += string(item)
	}
	res = append(res, d)
	return res
}

// удаление всех пробелов
func clearWhitespace(s string) string {
	res := ""
	for _, el := range s {
		if el == ' ' {
			continue
		}
		res += string(el)
	}
	return res
}

// определение класса елемента строки
func element(item int32) int {
	var res int
	switch {
	case '0' <= item && item <= '9':
		res = digit
	case item == '+' || item == '-':
		res = operation
	case item == '(' || item == ')':
		res = bracket
	}
	return res
}

func convertInDigit(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

/*


5+3-4 -(1+2-7+ (10-1 +3 +5+(3-0+(8-(3+(8-(10-(6-10-8-7+(0+0+7)-10+5-3-2+(9+0+(7+(2-(2-(9)-2+5+4+2+(2+9+1+5+5-8-9-2-9+1+0)-(5-(9)-(0-(7+9)+(10+(6-4+6))+0-2+(10+7+(8+(7-(8-(3)+(2)+(10-6+10-(2)-7-(2)+(3+(8))+(1-3-8)+6-(4+1)+(6))+6-(1)-(10+(4)+(8)+(5+(0))+(3-(6))-(9)-(4)+(2))))))-1)))+(9+6)+(0))))+3-(1))+(7))))))))
5 3 -4 -1 -2 7 -10 1 -3 -5 -3 0 -8  3  8 -10  6-10-8-7 0 0 7  -10 5-3-2 9 0 7 2 -2 9 2 -5 -4 -2 -2 -9 -1 -5 -5 8 9 2 9 -1 0 5 -9 0 7 9 -10 -6 4 -6 0 0 2 -10 -7 -8 -7 8 -3 2 10 -6 10 -2 -7 -2 3 8 0 1 -3 -8 6 -4 -1 6 0 6 -1 -10 -4 -8 -5 0 0 -3 6 0 -9 -4 2 0 -1 0 9 6 0 0 3 -1 0 -7 0]


*/
