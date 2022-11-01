/* https://leetcode.com/problems/jump-game/
55. Jump Game
Вам дается целочисленный массив nums. Изначально вы находитесь в первом индексе массива,
и каждый элемент в массиве представляет вашу максимальную длину перехода в этой позиции.
Верните значение true, если вы можете достичь последнего индекса, или false в противном случае.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Пояснение: Перейдите на 1 шаг от индекса 0 к 1, затем на 3 шага к последнему индексу.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Пояснение: Вы всегда придете к индексу 3, несмотря ни на что.
Его максимальная длина перехода равна 0, что делает невозможным достижение последнего индекса.

Выполнено:
Время выполнения: 127 мс, быстрее, чем 66,15% онлайн-заявок на игру Jump.
Использование памяти: 7,6 МБ, что составляет менее 54,41% от количества отправленных онлайн-заявок на игру Jump.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	//fileName := "050/55_jump_game/data.txt" // false
	//nums := readFile(fileName)

	//nums := []int{31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 0} // false
	//nums := []int{2, 3, 1, 1, 1, 2, 3, 1, 1, 2, 4} // true
	//nums := []int{2, 3, 1, 1, 4} // true
	//nums := []int{3, 2, 1, 0, 4} // false
	//nums := []int{1, 2, 0, 1} // true
	//nums := []int{2, 0} // true
	//nums := []int{2, 5, 0, 0} // true
	//nums := []int{0} // true
	//nums := []int{2, 0, 3, 1, 0, 0} //true
	//nums := []int{2, 0, 3} // true
	//nums := []int{1, 0, 3} // false
	//nums := []int{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6} // false
	//nums := []int{2, 0, 0, 9, 6} // false
	nums := []int{3, 0, 0, 9, 6} // true

	//fmt.Println("nums =", nums)
	start := time.Now()
	res := canJump(nums)
	fmt.Println("time  =", time.Now().Sub(start))
	fmt.Println("res  =", res)
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var target int
	for i := len(nums) - 2; i > -1; i-- {
		target++
		if nums[i] >= target {
			target = 0
			if i == 0 {
				return true
			}
		}
	}
	return false
}

// очень долго если элементов более 30
func canJumpRecursive(nums []int) bool {
	target := len(nums) - 1
	if target == 0 {
		return true
	}

	for i := 0; i < nums[0]; i++ {
		if nums[i] >= target {
			return true
		}
		if canJumpRecursive(nums[i+1:]) {
			return true
		}
	}
	return false
}

func readFile(fileName string) []int {
	f, errReadFile := ioutil.ReadFile(fileName)
	if errReadFile != nil {
		fmt.Println("error read file:", errReadFile.Error())
		return nil
	}
	if len(f) == 0 {
		fmt.Println("data no found")
		return nil
	}

	df := strings.Split(string(f), ",")
	data := make([]int, 0, len(df))
	for _, d := range df {
		d = strings.TrimSpace(d)
		i, errConv := strconv.Atoi(d)
		if errConv != nil {
			fmt.Println("error strconv:", errConv.Error())
			return nil
		}
		data = append(data, i)
	}

	return data
}
