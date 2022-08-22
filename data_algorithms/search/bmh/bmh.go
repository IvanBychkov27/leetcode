/* https://www.youtube.com/watch?v=kvWFAZwZ_8U&list=PLA0M1Bcd0w8yF0PO0eJ9v8VlsYEowmsnJ&index=2
Алгоритм Бойера-Мура-Хорспула
Поиск подстроки в строке
Сложность: max O(n*m) - сред O(n/|Z|) - min O(n/m), где |Z| - сумма неповторяющихся символов в строке и подстроке
*/
package main

import "fmt"

func main() {
	//a, t := "большие метеоданные", "данные"
	//a, t := "метеоданные", "данные"
	a, t := "test", "s"

	idx, ok := searchSubstringBMH(a, t)
	if ok {
		fmt.Printf("substring: %q - ok! - idx = %d\n", t, idx)
	} else {
		fmt.Printf("substring: %q - no found\n", t)
	}
}

// searchSubstringBMH - поиск подстроки в строке - Алгоритм Бойера-Мура-Хорспула - Сложность max O(n*m) - сред O(n/|Z|) - min O(n/m)
func searchSubstringBMH(a, t string) (int, bool) {
	m := len(t)
	if m == 0 || m > len(a) {
		return 0, false
	}

	// 1 этап - формирование таблицы смещения
	d := make(map[byte]int)
	for i := m - 2; i > -1; i-- {
		_, ok := d[t[i]]
		if !ok {
			d[t[i]] = m - i - 1
		}
	}
	_, ok := d[t[m-1]]
	if !ok {
		d[t[m-1]] = m
	}
	d['*'] = m

	// 2 этап - поиск подстроки
	n := len(a)
	i := m - 1
	for i < n {
		k := 0
		for j := m - 1; j > -1; j-- {
			if a[i-k] != t[j] {
				var off int
				if j == m-1 {
					off, ok = d[a[i]]
					if !ok {
						off = d['*']
					}
				} else {
					off = d[t[j]]
				}
				i += off
				break
			}
			k++
			if j == 0 {
				return i - k + 1, true
			}
		}
	}
	return 0, false
}
