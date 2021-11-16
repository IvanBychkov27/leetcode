// https://leetcode.com/problems/palindrome-number/
/*
Учитывая целое число x, верните значение true, если x является целым числом палиндрома.
Целое число является палиндромом, когда оно читается одинаково как вперед, так и назад. Например, 121 - это палиндром, а 123 - нет.

Результат:
Время выполнения: 8 мс, быстрее, чем 92,93 % отправленных онлайн-заявок на номер палиндрома.
Использование памяти: 5,1 МБ, менее 87,93 % отправленных онлайн-заявок на номер палиндрома.
*/

package main

import (
	"fmt"
)

func main() {
	x := 1101011
	res := isPalindrome(x)
	fmt.Println("res  =", res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	data := make([]int, 0, 10)
	for {
		d := x % 10
		data = append(data, d)
		x /= 10
		if x == 0 {
			break
		}
	}

	l := len(data)
	if l > 10 {
		return false
	}

	for i := 0; i < l/2; i++ {
		if data[i] != data[l-1-i] {
			return false
		}
	}

	return true
}
