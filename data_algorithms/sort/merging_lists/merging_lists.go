/* https://www.youtube.com/watch?v=GqieI_1yh40&list=PLA0M1Bcd0w8yF0PO0eJ9v8VlsYEowmsnJ&index=11
Метод слияния двух отсортированных списков (массивов)
*/

package main

import "fmt"

func main() {
	a := []int{1, 4, 10, 11}
	b := []int{2, 3, 3, 3, 4, 8}

	res := mergingLists(a, b)
	fmt.Println(res)
}

// mergingLists - метод слияния двух отсортированных списков (массивов)
func mergingLists(a, b []int) []int {
	n := len(a)
	m := len(b)
	c := make([]int, 0, n+m)

	i := 0
	j := 0

	for i < n && j < m {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	c = append(c, a[i:]...)
	c = append(c, b[j:]...)

	return c
}
