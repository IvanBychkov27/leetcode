/* https://leetcode.com/problems/valid-parentheses/
Дана строка s, содержащая только символы '(', ')', '{', '}', '[' и "]", определите, является ли входная строка допустимой.
Входная строка допустима, если:
Открытые скобки должны быть закрыты скобками того же типа.
Открытые скобки должны быть закрыты в правильном порядке.

Input: s = "{[]}"
Output: true

Input: s = "([)]"
Output: false

Решено:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправлений в режиме онлайн для допустимых скобок.
Использование памяти: 2 МБ, менее 100,00 % отправленных онлайн-сообщений для допустимых скобок.

*/
package main

import "fmt"

func main() {
	//s := "{([])}" // true
	//s := "{([)]}" // false
	//s := "{[]}" // true
	s := "(){}[]"

	res := isValid(s)
	fmt.Println("res =", res)
}

func isValid(s string) bool {
	pred := make([]byte, 0, len(s))

	sim := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
		'{': '.',
		'(': '.',
		'[': '.',
	}

	data := []byte(s)
	for _, d := range data {
		v, ok := sim[d]
		if !ok {
			continue
		}
		if v == '.' {
			pred = append(pred, d)
			continue
		}
		if len(pred) == 0 || pred[len(pred)-1] != sim[d] {
			return false
		}
		pred = pred[:len(pred)-1]
	}

	if len(pred) > 0 {
		return false
	}

	return true
}
