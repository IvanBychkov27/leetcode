/*
В строке s, найдите длину самой длинной подстроки без повторяющихся символов

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
	s := "pwwkewwacd" // 4

	res := lengthOfLongestSubstring(s)
	fmt.Println("res =", res)
}

func lengthOfLongestSubstring(s string) int {
	var idx, res int
	sRes := make([]byte, 0, len(s))
	temp := make([]byte, 0, len(s))
	m := make(map[byte]struct{})
	data := []byte(s)
	for j := 0; j < len(data)-res; j++ {
		for i := j; i < len(data); i++ {
			d := data[i]
			_, ok := m[d]
			if ok {
				if res < idx {
					res = idx
					sRes = temp
				}
				m = make(map[byte]struct{})
				idx = 0
				temp = nil
				continue
			}
			m[d] = struct{}{}
			temp = append(temp, d)
			idx++
		}

		if res < idx {
			res = idx
			sRes = temp
		}
		m = make(map[byte]struct{})
		idx = 0
		temp = nil

	}
	fmt.Println("str =", string(sRes))
	return res
}
