/* https://leetcode.com/problems/longest-valid-parentheses/

Учитывая строку, содержащую только символы '(' и ')',
найдите длину самой длинной допустимой (правильно сформированной) подстроки в круглых скобках.

*/
package main

import "fmt"

func main() {
	//s := "(()" // 2
	//s := "())" // 2
	//s := "()()" // 4
	//s := "" // 0
	//s := ")(" // 0
	//s := ")()())" // 4
	//s := "()()())"  // 6
	//s := ")(()()())" // 8
	//s := "()(((())((()))))" // 2
	//s := "(()(()" // 2
	//s := "())())" // 2
	s := ")))()(()))())(())()())(()((())))())))))(())))(()()))(())())())))(()))()))((((()())))))()()))(()((())((())())()()()()((()((((())))(()))(()()()))))(()((((()))(((((()))())()))(((" // 52

	res := longestValidParentheses(s)
	fmt.Println("res =", res)
}

func longestValidParentheses(s string) int {
	res := process(s)
	m := process(reverse(s))

	if m < res {
		res = m
	}
	return res
}

func process(s string) int {
	max := 0
	count := 0
	n := 0
	pred := 0
	for _, d := range s {
		count++
		if d == '(' {
			n++
		}
		if d == ')' {
			n--
		}

		if n == -1 {
			if max < count-1 {
				max = count - 1
			}
			count = 0
			n = 0
			pred = 0
			continue
		}

		if n == 0 {
			pred += count
			if max < pred {
				max = pred
			}
			count = 0
		}

	}

	count -= n
	if max < count {
		max = count
	}
	if max < pred {
		max = pred
	}
	return max
}

func reverse(s string) string {
	rev := ""
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '(' {
			rev += ")"
		} else {
			rev += "("
		}
	}
	return rev
}
