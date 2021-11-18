// https://leetcode.com/problems/regular-expression-matching/
/*
Учитывая входную строку s и шаблон p, реализуйте сопоставление регулярных выражений с поддержкой '.' и '*', где:
"." Соответствует любому отдельному символу.
'*' Соответствует нулю или более предыдущего элемента.
Совпадение должно охватывать всю входную строку (не частично).

Input: s = "aa", p = "a"
Output: false
Explanation: "a" не соответствует всей строке "aa".

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' означает ноль или более предыдущего элемента, 'a'. Поэтому, повторяя "а" один раз, оно становится "аа".

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c - м.б. повторено 0 раз, а - м.б. повторено 1 раз. Следовательно, соответствует "aab".

Input: s = "mississippi", p = "mis*is*p*."
Output: false
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	//s, p := "aab", "c*a*b" // true
	//s, p := "ab", ".*" // true
	//s, p := "aa", "a*" // true
	//s, p := "aa", "a" // false
	//s, p := "ab", ".*c" // false
	//s, p := "aaa", "ab*a*c*a" // true
	//s, p := "mississipp", "mis*is*ip*" // true
	//s, p := "mississippi", "mis*is*ip*." // true
	//s, p := "i", "." // true
	//s, p := "aaa", "aaaa" // false
	s, p := "aaba", "ab*a*c*a" // false
	//s, p := "aaca", "ab*a*c*a" // true

	res := isMatch(s, p)
	fmt.Println("res  =", res)
}

func isMatch(s string, p string) bool {
	var idxD, idxS int
	var predD, predS byte
	data := []byte(s)
	shablon := []byte(p)
	lenD := len(data)
	lenS := len(shablon)

	if !bytes.Contains(shablon, []byte(".")) && !bytes.Contains(shablon, []byte("*")) {
		return s == p
	}

	for {
		if idxD >= lenD {
			if data[lenD-1] != shablon[lenS-1] && shablon[lenS-1] != '*' && shablon[lenS-1] != '.' {
				return false
			}
			break
		}
		if idxS >= lenS {
			return false
		}
		if idxD > 0 {
			predD = data[idxD-1]
		}
		if idxS > 0 {
			predS = shablon[idxS-1]
		}

		if data[idxD] != shablon[idxS] && shablon[idxS] == '.' {
			idxD++
			idxS++
			continue
		}

		if data[idxD] != shablon[idxS] && shablon[idxS] != '.' && shablon[idxS] != '*' && idxS < lenS-1 && shablon[idxS+1] != '*' {
			return false
		}

		if data[idxD] == shablon[idxS] {
			idxD++
			idxS++
			continue
		}

		if data[idxD] != shablon[idxS] && idxS < lenS-1 && shablon[idxS+1] == '*' && shablon[idxS] != '.' {
			idxS += 2
			continue
		}

		if data[idxD] != shablon[idxS] && shablon[idxS] == '*' {
			if data[idxD] != predD && predS != '.' && predD != predS {
				return false
			}

			if data[idxD] != predD && predD == predS {
				idxS++
				continue
			}

			idxD++
			idxS++
			continue
		}
		idxD++
		idxS++
	}
	return true
}
