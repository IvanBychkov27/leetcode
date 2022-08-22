package main

import "fmt"

func main() {
	//n := 20
	//numsFib := numbersFibonacciInArray(n)
	//fmt.Println("numbers Fibonacci:", numsFib)

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
