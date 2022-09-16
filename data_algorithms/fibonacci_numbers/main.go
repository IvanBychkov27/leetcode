package main

import "fmt"

func main() {
	n := 40
	numbersFibonacciLine(n)
	fmt.Printf("\n\n")

	resR := fibRecursive(n)
	fmt.Printf("resR = %d\n\n", resR)

	resD := fibDynamicArray(n)
	fmt.Printf("resD = %d\n\n", resD)

	resABC := fibDynamicABC(n)
	fmt.Printf("resABC = %d\n\n", resABC)

	numsFib := numbersFibonacciInArray(n)
	fmt.Println("numbers Fibonacci:", numsFib)

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	target := 12
	fmt.Println(searchFibonacci(nums, target))

}

// Поиск Фибоначчи - это вариант бинарного поиска. Его временная сложность O (log2n).
func searchFibonacci(nums []int, target int) bool {
	if nums[0] > target || nums[len(nums)-1] < target {
		return false
	}

	var a, b, idx int
	if nums[idx] == target {
		return true
	}
	a = 1
	for {
		idx = a + b

		if idx >= len(nums) {
			return searchFibonacci(nums[b:], target)
		}

		if nums[idx] == target {
			return true
		}

		if nums[idx] > target {
			return searchFibonacci(nums[b:idx], target)
		}

		a = b
		b = idx
	}
}

// динамическое программирование
func numbersFibonacciLine(n int) {
	var a, b int
	for i := 0; i < n; i++ {
		if i == 0 {
			a = 1
		}
		c := a + b
		fmt.Print(c, " ")
		a = b
		b = c
	}
}

// динамическое программирование
func numbersFibonacciInArray(n int) []int {
	if n < 1 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 1}
	}

	nums := make([]int, 0, n)
	nums = append(nums, []int{1, 1}...)
	for i := 2; i < n; i++ {
		f := nums[i-1] + nums[i-2]
		nums = append(nums, f)
	}

	return nums
}

// Рекурсия работает долго!!! - время работы растет прямо пропорционально числам фибоначчи (введите например n=45 и более)
// Сложность O(fib(n))
func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// динамическое программирование вычисление числа фибоначчи
// Сложность O(n)
func fibDynamicArray(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func fibDynamicABC(n int) int {
	var a, b, c int
	b = 1
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
