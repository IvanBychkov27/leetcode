/* https://www.youtube.com/watch?v=5BuCMzKYagg&list=PLA0M1Bcd0w8yF0PO0eJ9v8VlsYEowmsnJ&index=10
	Сортировка пузырьком
	Сложность O(n^2)
Примечание:
Сложность сортировки пузырьковым методом составляет O (n2).
На практике данный метод почти не используется, вместо него применяются более эффективные алгоритмы сортировки данных.
Сортировка выбором для массива из n элементов имеет время выполнения в худшем, среднем и лучшем случае O (n2),
предполагая что сравнения делаются за постоянное время.
*/

package main

import (
	"fmt"
	"time"
)

//var count int64

func main() {
	a := setData(20)
	fmt.Println(a)

	timeStart := time.Now()
	sortA := sortingBubble(a)
	fmt.Println(time.Now().Sub(timeStart))
	_ = sortA
	fmt.Println(sortA)
	//fmt.Println("count:", count)
}

// sortingBubble - сортировка пузырьком
func sortingBubble(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			//atomic.AddInt64(&count, 1)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func setData(n int) []int {
	a := make([]int, 0, n)
	d := n / 2
	for i := 0; i < n; i++ {
		a = append(a, d-i)
	}
	return a
}
