/* https://www.youtube.com/watch?v=S2I0covkyMc&list=PLA0M1Bcd0w8yF0PO0eJ9v8VlsYEowmsnJ&index=1
Алгоритм Кнута-Морриса-Пратта (КМП-алгоритм)
Поиск подстроки в строке
Сложность O(n+m)
*/
package main

import "fmt"

func main() {
	//a, t := "лилилось лилилась", "лилила"
	a, t := "test", "s"

	idx, ok := searchSubstringKMP(a, t)
	if ok {
		fmt.Printf("substring: %q - ok! - idx = %d\n", t, idx)
	} else {
		fmt.Printf("substring: %q - no found\n", t)
	}

}

// searchSubstringKMP - поиск подстроки в строке - Алгоритм Кнута-Морриса-Пратта (КМП-алгоритм) - Сложность O(m+n)
func searchSubstringKMP(a, t string) (int, bool) {
	m := len(t)
	if m == 0 {
		return 0, false
	}

	// 1 - этап заполнения вспомогательного массива p
	p := make([]int, m)
	i := 1
	j := 0
	for i < m {
		if t[j] == t[i] {
			p[i] = j + 1
			i++
			j++
			continue
		}
		if j == 0 {
			p[i] = 0
			i++
		} else {
			j = p[j-1]
		}
	}

	// 2 - основной этап поиска подстроки в строке
	n := len(a)
	i = 0
	j = 0
	for i < n {
		if a[i] == t[j] {
			i++
			j++
			if j == m {
				return i - m, true
			}
			continue
		}
		if j > 0 {
			j = p[j-1]
		} else {
			i++
		}
	}

	return 0, false
}
