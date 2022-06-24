// https://nuancesprog.ru/p/12140/?ysclid=l4sa0j6nde701881476
/*
Основные операции с массивами

    Get — получить элемент массива по заданному индексу.
    Insert — вставить элемент массива по заданному индексу.
    Length — получить количество элементов в заданном массиве.
    Delete — удалить элемент массива по заданному индексу. Может быть выполнено либо путем установки значения undefined, либо путем копирования элементов массива, за исключением удаляемого, в новый массив.
    Update — обновление значения элемента массива по заданному индексу.
    Traverse — проход цикла через массив для выполнения функций над элементами массива.
    Search — поиск определенного элемента в заданном массиве с помощью выбранного алгоритма.
	Sort — сортировка массива по возрастанию с помощью выбранного алгоритма.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	//data := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	data := []int{90, 80, 70, 60, 50, 40, 30, 20, 10, 0}

	idx := 5
	value, err := Get(idx, data)
	checkErr(err)
	fmt.Printf("data[%d] = %d \n", idx, value)

	d := 55
	data = Insert(idx, d, data)
	printInfo("Insert idx", idx, data)

	data, err = Delete(idx, data)
	checkErr(err)
	printInfo("Delete idx", idx, data)

	idx = 0
	data, err = Delete(idx, data)
	checkErr(err)
	printInfo("Delete idx", idx, data)

	idx = len(data) - 1
	data, err = Delete(idx, data)
	checkErr(err)
	printInfo("Delete idx", idx, data)

	idx = 2
	d = 33
	data, err = Update(idx, d, data)
	checkErr(err)
	printInfo("Update idx", idx, data)

	d = 5
	data = Traverse(data, d, addArray)
	printInfo("Traverse add  ", d, data)

	d = 3
	data = Traverse(data, d, minusArray)
	printInfo("Traverse minus", d, data)

	d = 10
	data = Traverse(data, d, multArray)
	printInfo("Traverse mult", d, data)

	d = 2
	data = Traverse(data, d, delArray)
	printInfo("Traverse del  ", d, data)

	d = 410
	ok, index := Search(data, d, searchInIntSlice)
	if ok {
		fmt.Printf("Search: the desired element %d is found - idx = %d \n", d, index)
	} else {
		fmt.Printf("Search: the desired element %d was not found \n", d)
	}

	data = Sort(data, quickSort)
	printInfo("Sort quickSort", 0, data)

}

func printInfo(msg string, idx int, data []int) {
	fmt.Printf("%s %d - data: %v len = %d\n", msg, idx, data, Length(data))

}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

// Get - получить элемент массива по заданному индексу
func Get(idx int, data []int) (int, error) {
	if idx >= len(data) {
		return 0, fmt.Errorf("the index %d exceeded the length of the array", idx)
	}
	return data[idx], nil
}

// Insert — вставить элемент массива по заданному индексу
func Insert(idx, d int, data []int) []int {
	newArray := make([]int, 0, len(data)+1)
	newArray = append(newArray, data[:idx]...)
	newArray = append(newArray, d)
	newArray = append(newArray, data[idx:]...)
	return newArray
}

// Length - получить количество элементов в заданном массиве
func Length(data []int) int {
	return len(data)
}

// Delete — удалить элемент массива по заданному индексу
func Delete(idx int, data []int) ([]int, error) {
	if idx >= len(data) {
		return nil, fmt.Errorf("the index %d exceeded the length of the array", idx)
	}
	newArray := make([]int, 0, len(data)-1)
	newArray = append(newArray, data[:idx]...)
	if idx != len(data) {
		newArray = append(newArray, data[idx+1:]...)
	}
	return newArray, nil
}

// Updete — обновление значения элемента массива по заданному индексу
func Update(idx, d int, data []int) ([]int, error) {
	if idx >= len(data) {
		return nil, fmt.Errorf("the index %d exceeded the length of the array", idx)
	}
	data[idx] = d
	return data, nil
}

//  Traverse — проход цикла через массив для выполнения функций над элементами массива
func Traverse(data []int, d int, f func(a, b int) int) []int {
	for idx, val := range data {
		data[idx] = f(val, d)
	}
	return data
}

func addArray(a, b int) int {
	return a + b
}

func minusArray(a, b int) int {
	return a - b
}

func multArray(a, b int) int {
	return a * b
}

func delArray(a, b int) int {
	return int(a / b)
}

//  Search — поиск определенного элемента в заданном массиве с помощью выбранного алгоритма
func Search(data []int, d int, f func(a []int, b int) (bool, int)) (bool, int) {
	return f(data, d)
}

// binary_search
func searchInIntSlice(data []int, n int) (res bool, idx int) {
	sort.Ints(data)
	lowKey := 0              // первый индекс
	highKey := len(data) - 1 // последний индекс
	if (data[lowKey] > n) || (data[highKey] < n) {
		return // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		// уменьшаем список рекурсивно
		idx = (lowKey + highKey) / 2 // середина
		if data[idx] == n {
			res = true // мы нашли значение
			return
		}
		if data[idx] < n {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = idx + 1
			continue
		}
		if data[idx] > n {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = idx - 1
		}
	}
	return
}

// Sort — сортировка массива по возрастанию с помощью выбранного алгоритма
func Sort(data []int, f func(a []int) []int) []int {
	return f(data)
}

// 	Встроенная в Go сортировка
func goSort(data []int) []int {
	sort.Ints(data)
	return data
}

// 	Быстрая сортировка
func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	less := make([]int, 0)
	bigger := make([]int, 0)
	n := data[0]
	for _, v := range data[1:] {
		if v > n {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	data = append(quickSort(less), n)
	data = append(data, quickSort(bigger)...)
	return data
}
