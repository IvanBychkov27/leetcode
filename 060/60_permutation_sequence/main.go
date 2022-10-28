/* https://leetcode.com/problems/permutation-sequence/
60. Permutation Sequence - Последовательность перестановок

Множество [1, 2, 3, ..., n] содержит в общей сложности n! уникальных перестановок.
Перечисляя и помечая все перестановки по порядку, мы получаем следующую последовательность для n = 3:
"123", "132", "213", "231", "312", "321"
Учитывая n и k, верните k-ю последовательность перестановок.

Example 1:
Input: n = 3, k = 3
Output: "213"

Example 2:
Input: n = 4, k = 9
Output: "2314"

Example 3:
Input: n = 3, k = 1
Output: "123"


*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//n, k := 3, 1 // 123
	//n, k := 4, 9 // 2314
	//n, k := 3, 3 // 213
	n, k := 3, 5 // "312"
	//n, k := 8, 31492 // 72641583
	//n, k := 9, 331987 // 928157346
	//n, k := 9, 136371 // 451672839

	res := getPermutation(n, k)
	fmt.Println("res =", res)
}

//var dd [][]int
var dd []string

func getPermutation(n int, k int) string {
	d := make([]int, 0, n)
	for i := 1; i < n+1; i++ {
		d = append(d, i)
	}
	dd = make([]string, 0, k)
	permutationsInArray(d, 0)

	//res := convertInString(k)
	//return res[k-1]

	fmt.Println(dd)
	//sort.Strings(dd)
	return dd[k-1]
}

// перестановки сохраняем в массив
func permutationsInArray(d []int, n int) {
	l := len(d)
	if n == l-1 {
		a := ""
		for _, t := range d {
			a += strconv.Itoa(t)
		}
		dd = append(dd, a)
		return
	}
	for i := n; i < l; i++ {
		d[n], d[i] = d[i], d[n]
		permutationsInArray(d, n+1)
		d[n], d[i] = d[i], d[n]
	}
}

//
//func convertInString(k int) []string {
//	data := make([]string, 0, len(dd))
//	k += 50000
//	for _, d := range dd {
//		if k == 0 {
//			break
//		}
//		s := ""
//		for _, el := range d {
//			s += strconv.Itoa(el)
//		}
//		data = append(data, s)
//		k--
//	}
//	sort.Strings(data)
//	return data
//}
