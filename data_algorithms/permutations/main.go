// перестановки без повторений

package main

import "fmt"

var dd [][]int

func main() {
	d := []int{1, 2, 3}
	f := factorial(len(d))
	fmt.Printf("f(%d) = %d\n", len(d), f)

	permutations(d, 0)
	fmt.Println()

	dd = make([][]int, 0)
	permutationsInArray(d, 0)

	fmt.Println(dd)

}

// перестановки сохраняем в массив
func permutationsInArray(d []int, n int) {
	if n == len(d)-1 {
		a := make([]int, 0, len(d))
		for _, t := range d {
			a = append(a, t)
		}
		dd = append(dd, a)
		return
	}

	for i := n; i < len(d); i++ {
		d[n], d[i] = d[i], d[n]
		permutationsInArray(d, n+1)
		d[n], d[i] = d[i], d[n]
	}
}

// перестановки выводим на экран
func permutations(d []int, n int) {
	if n == len(d)-1 {
		fmt.Println(d)
		return
	}

	for i := n; i < len(d); i++ {
		d[n], d[i] = d[i], d[n]
		permutations(d, n+1)
		d[n], d[i] = d[i], d[n]
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
