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

	fmt.Println()
	fmt.Println("factorialOK =", factorialOK(5))

	fmt.Println()
	n := 3
	comb := make([]int, n)
	generateBinaryNumber(comb, n, 0)
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

// факториал через рекурсию считать неправильно!!! (можно переполнить стек памяти, при больших числах)
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// рачсет факториала в цикле
func factorialOK(n int) int {
	f := 1
	if n == 0 {
		return f
	}
	for i := 1; i < n+1; i++ {
		f *= i
	}
	return f
}

// генерация всех вариантов двоичного числа длиной n
func generateBinaryNumber(comb []int, n, idx int) {
	if n == 0 {
		fmt.Println(comb)
		return
	}

	comb[idx] = 0
	idx++
	generateBinaryNumber(comb, n-1, idx)
	idx--

	comb[idx] = 1
	idx++
	generateBinaryNumber(comb, n-1, idx)
	idx--
}
