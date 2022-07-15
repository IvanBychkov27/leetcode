package main

import (
	"fmt"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	type data struct {
		s string
		r int
	}
	ds := []data{
		{"(()", 2},
		{"(()", 2},
		{"())", 2},
		{"()()", 4},
		{"", 0},
		{")(", 0},
		{")()())", 4},
		{"()()())", 6},
		{"()((()))", 8},
		{"))(()()())((()))((", 14},
		{"()(()", 2},
		{"(()(((()", 2},
		{")))()(()))())(())()())(()((())))())))))(())))(()()))(())())())))(()))()))((((()())))))()()))(()((())((())())()()()()((()((((())))(()))(()()()))))(()((((()))(((((()))())()))(((", 52},
	}
	var res int
	for _, d := range ds {
		res = longestValidParentheses(d.s)
		if res != d.r {
			fmt.Printf("Not equal:\n s = %s \nexpected: %d\nactual  : %d\n", d.s, d.r, res)
			fmt.Println()
		}
	}

}
