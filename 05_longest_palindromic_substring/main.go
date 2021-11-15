// https://leetcode.com/problems/longest-palindromic-substring/
/*
Учитывая строку s, верните самую длинную палиндромную подстроку в s.

Результат:
Время выполнения: 152 мс, быстрее, чем 31,64 % онлайн-отправлений для самой длинной палиндромной подстроки.
Использование памяти: 2,9 МБ, менее 44,41 % отправленных онлайн-сообщений для самой длинной палиндромной подстроки.
*/

package main

import (
	"fmt"
)

func main() {
	s := "babad" // res = "bab"
	//s := "cbbd" // res = "bb"
	//s := "a" // res = "a"
	//s := "ac" // res = "a"
	//s := "bab" // res = "a"

	res := longestPalindrome(s)
	fmt.Println("res =", res)
}

func longestPalindrome(s string) string {
	var maxLen int
	var p bool
	var res []byte

	data := []byte(s)
	lenData := len(data)
	for j := 0; j < lenData; j++ {
		for i := j; i < lenData; i++ {
			d := data[j : i+1]
			l := len(d)
			p = true
			for idx := 0; idx < l/2; idx++ {
				if d[idx] != d[l-1-idx] {
					p = false
					break
				}
			}
			if p && maxLen < l {
				maxLen = l
				res = d
			}
		}
		if (j + maxLen) >= lenData {
			break
		}
	}
	return string(res)
}

func pol(s string) string {
	d := []byte(s)
	l := len(d)
	for idx := 0; idx < l/2; idx++ {
		if d[idx] != d[l-1-idx] {
			return "no"
		}
	}
	return "ok"
}
