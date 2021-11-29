//https://leetcode.com/problems/container-with-most-water/
/*
Container With Most Water - Контейнер С Наибольшим Количеством Воды
Дано n неотрицательных целых чисел a1, a2,..., an, где каждое представляет точку с координатой (i, ai).
n вертикальных линий нарисованы таким образом, что две конечные точки линии i находятся в (i, ai) и (i, 0).
Найдите две линии, которые вместе с осью x образуют контейнер, такой, чтобы в контейнере было больше всего воды.
Обратите внимание, что вы не можете наклонять контейнер.

Входные данные: высота = [1,8,6,2,5,4,8,3,7]
Выход: 49
Пояснение: Вышеуказанные вертикальные линии представлены массивом [1,8,6,2,5,4,8,3,7].
		   В этом случае максимальная площадь воды (синяя секция), которую может содержать контейнер, составляет 49.
*/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	//height, expected := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49
	//height, expected  := []int{1, 1}, 1
	//height, expected  := []int{4, 3, 2, 1, 4}, 16
	//height, expected  := []int{1, 2, 1}, 2
	//height, expected := []int{3000, 8000, 6, 2, 7, 5, 4, 8, 3, 7, 101, 100}, 3000
	height, expected := []int{30, 80, 6, 2, 7, 5, 4000, 8, 3000, 7, 101, 100}, 6000

	//height, expected := getBigData("11_container_with_most_water/bigdata_705634720.txt"), 705634720 // len = 71050
	//height, expected := getBigData("11_container_with_most_water/bigdata_402471897.txt"), 402471897 // len = 40566

	fmt.Println("n =", len(height))

	timeStart := time.Now()
	res := maxArea(height)
	timeEnd := time.Now()

	fmt.Println("time:", timeEnd.Sub(timeStart))
	fmt.Println()

	fmt.Println("res =", res)
	if res == expected {
		fmt.Println("OK")
	} else {
		fmt.Println("NO")
	}

	fmt.Println("res_01 =", maxArea_01(height))
	fmt.Println("res_02 =", maxArea_02(height))
}

func maxArea(height []int) int {
	var max, i, j, im int
	f := false
	j = len(height) - 1
	for {
		if i == j && f {
			break
		}

		min := height[i]
		if min > height[j] {
			min = height[j]
		}
		val := min * (j - i)

		if max < val {
			if !f {
				im = i
			}
			max = val
		}

		if f {
			j--
		} else {
			i++
		}

		if i == j {
			f = true
			i = im
		}

	}

	return max
}

func maxArea_02(height []int) int {
	var res int
	ch := make(chan int, len(height))

	for jj := 0; jj < len(height); jj++ {
		go func(j int, ch chan int, height []int) {
			var max int
			for i := j + 1; i < len(height); i++ {
				min := height[j]
				if min > height[i] {
					min = height[i]
				}

				val := min * (i - j)

				if max < val {
					max = val
				}
			}
			ch <- max
		}(jj, ch, height)
	}

	for i := 0; i < len(height); i++ {
		d := <-ch
		if res < d {
			res = d
		}
	}

	return res
}

func maxArea_01(height []int) int {
	var max int
	for j := 0; j < len(height); j++ {
		for i := j + 1; i < len(height); i++ {
			min := height[j]
			if min > height[i] {
				min = height[i]
			}

			val := min * (i - j)

			if max < val {
				max = val
			}
		}
	}
	return max
}

func getBigData(fileName string) []int {
	df, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error:", err.Error())
		return nil
	}
	data := bytes.Split(df, []byte(","))

	res := make([]int, 0, len(data))
	for _, b := range data {
		d, errConv := strconv.Atoi(string(bytes.TrimSpace(b)))
		if errConv != nil {
			fmt.Println("error strconv:", errConv.Error())
			return nil
		}
		res = append(res, d)
	}
	return res
}
