/* https://leetcode.com/problems/length-of-last-word/
58. Length of Last Word - Длина последнего слова

Учитывая строку s, состоящую из слов и пробелов, верните длину последнего слова в строке.
Слово - это максимальная подстрока, состоящая только из символов без пробелов.

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00% онлайн-заявок на длину последнего слова.
Использование памяти: 2 МБ, менее 100,00% отправленных онлайн-заявок из-за длины последнего слова.

*/

package main

import "fmt"

func main() {
	//s := "Hello world" // 5
	s := "   fly me   to   the moon  " // 4
	//s := "luffy is still joyboy" // 6
	//s := "" // 0

	res := lengthOfLastWord(s)
	fmt.Println("res =", res)
}

func lengthOfLastWord(s string) int {
	var count int
	lenS := len(s)
	for idx := lenS - 1; idx > -1; idx-- {
		if s[idx] == 32 {
			if count == 0 {
				continue
			} else {
				break
			}
		}
		count++
	}

	return count
}
