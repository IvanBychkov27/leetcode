/* https://algorithmica.org/ru/karatsuba?ysclid=l7q9re4l5n748788033


 */
package main

import (
	"fmt"
)

func main() {
	asA := "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	bsA := "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	arrayASA := parseDataInArray(asA)
	arrayBSA := parseDataInArray(bsA)
	resSA := multiplyArray(arrayASA, arrayBSA)
	fmt.Println("asA * bsA = ", printRes(resSA))
	fmt.Println("len = ", len(resSA))
}

func parseDataInArray(num string) []int {
	const correct = 48
	var digit int
	n := len(num)
	res := make([]int, 0, n)
	for i := n - 1; i > -1; i-- {
		digit = int(num[i] - correct)
		if i == 0 && digit == 0 {
			continue
		}
		res = append(res, digit)
	}
	return res
}

// умножение цифр массивов - наивный алгоритм O(n^2)
func multiplyArray(a, b []int) []int {
	const base = 10
	n := len(a)
	m := len(b)

	res := make([]int, n+m)
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			d := a[i] * b[j]
			res[i+j] += d % base
			res[i+j+1] += d / base

			if res[i+j] >= base {
				res[i+j], res[i+j+1] = res[i+j]%base, res[i+j+1]+(res[i+j]/base)
			}
		}
	}
	return res
}

func printRes(d []int) string {
	var s string
	f := true
	for i := len(d) - 1; i > -1; i-- {
		if f {
			if d[i] == 0 {
				continue
			}
			f = false
		}
		s += fmt.Sprintf("%d", d[i])
	}
	if s == "" {
		s = "0"
	}
	return s
}
