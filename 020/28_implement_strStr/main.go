/* https://leetcode.com/problems/implement-strstr/

Реализуйте strStr().
Учитывая две строки needle и haystack, верните индекс первого появления иглы в стоге сена или -1, если игла не является частью стога сена.

Уточнение: Что мы должны вернуть, если игла - это пустая строка? Это отличный вопрос, который можно задать во время собеседования.
Для целей этой задачи мы вернем 0, если needle является пустой строкой. Это согласуется с strstr() языка Си и индексом Java Of().

Выполнено:
Время выполнения: 2 мс, быстрее, чем 40,30% отправленных онлайн-заявок для реализации strStr().
Использование памяти: 1,9 МБ, менее 67,65% отправленных онлайн-заявок для реализации strStr().

*/

package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "el"

	idx := strStr(haystack, needle)
	fmt.Println("idx =", idx)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n := len(needle)
	for idx := 0; idx < len(haystack)-n+1; idx++ {
		if haystack[idx:idx+n] == needle {
			return idx
		}
	}
	return -1
}
