/* https://leetcode.com/problems/longest-common-prefix/
Напишите функцию, чтобы найти самую длинную общую строку префикса среди массива строк.
Если общего префикса нет, верните пустую строку "".

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправлений в режиме онлайн для самого длинного общего префикса.
Использование памяти: 2,4 МБ, менее 100,00 % отправленных онлайн-сообщений для самого длинного общего префикса.
*/
package main

import "fmt"

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	strs := []string{"abc", "abc", "abc"}

	res := longestCommonPrefix(strs)

	fmt.Println("res =", res)

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	d := strs[0]
	if len(d) == 0 {
		return ""
	}
	max := len(d)
	for i := 1; i < len(strs); i++ {
		e := strs[i]
		if len(e) == 0 {
			return ""
		}
		idx := 0
		m := 0
		for {
			if d[idx] == e[idx] {
				m++
			} else {
				if max > m {
					max = m
				}
				break
			}
			idx++
			if idx == len(d) || idx == len(e) {
				if max > m {
					max = m
				}
				break
			}
		}
		if max == 0 {
			return ""
		}
	}

	return d[:max]
}
