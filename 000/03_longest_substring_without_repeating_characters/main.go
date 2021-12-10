/*
В строке s, найдите длину самой длинной подстроки без повторяющихся символов

Решено:
Время выполнения: 328 мс, быстрее, чем 12,67 % отправленных онлайн-сообщений для самой длинной подстроки Без повторяющихся символов.
Использование памяти: 7,1 МБ, менее 18,57 % отправленных онлайн-сообщений для самой длинной подстроки Без повторяющихся символов.
*/
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import (
	"fmt"
)

func main() {
	//s := "abcabcbb" // 3
	//s := "bbbbb" // 1
	//s := "" // 0
	//s := "pwwkew" // 3
	//s := "pwwkewwacd" // 4
	s := "pwwkewwacwkdfg" // 7

	res := lengthOfLongestSubstring(s)
	fmt.Println("res =", res)
}

func lengthOfLongestSubstring(s string) int {
	var idx, res, p int
	m := make(map[byte]int)
	data := []byte(s)
	for j := 0; j < len(data)-res; j++ {
		for i := j; i < len(data); i++ {
			var ok bool
			d := data[i]
			p, ok = m[d]
			if ok {
				j = p
				break
			}
			m[d] = i
			idx++
		}

		if res < idx {
			res = idx
		}
		m = make(map[byte]int)
		idx = 0
	}
	return res
}
