/* https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Удалить N-Й узел Из Конца списка
Учитывая заголовок связанного списка, удалите n-й узел из конца списка и верните его заголовок.

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Input: head = [1,2], n = 1
Output: [1]

Input: head = [1], n = 1
Output: []

Решено:
Время выполнения: 2 мс, быстрее, чем 16,77 % отправленных онлайн-сообщений для удаления N-го узла из конца списка.
Использование памяти: 2,6 МБ, менее 12,26 % отправленных онлайн сообщений для удаления N-го узла из конца списка.

*/

package main

import "fmt"

// определение односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//d, n := []int{5, 4, 3, 2, 1}, 1 // [1,2,3,5]
	//d, n := []int{1, 2, 3, 4, 5}, 2 // [1,2,3,5]
	//d, n := []int{1, 2, 3, 4, 5}, 1 // [1,2,3,5]
	//d, n := []int{1}, 1 // []
	d, n := []int{2, 1}, 1 // [1]
	head := NewList(d)
	fmt.Println("head =", getDataList(head))

	res := removeNthFromEnd(head, n)
	fmt.Println("res =", getDataList(res))

	fmt.Println()
	fmt.Println("new =", getDataList(removeNthFromEnd2(head, n)))

}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var lenHead int // длина списка
	var res, listRes *ListNode
	temp := head
	for {
		lenHead++
		res = &ListNode{Val: temp.Val, Next: res}
		temp = temp.Next
		if temp == nil {
			break
		}
	}

	temp = res
	step := 0
	for {
		step++
		if step == n {
			temp = temp.Next
			if temp == nil {
				break
			}
		}
		listRes = &ListNode{Val: temp.Val, Next: listRes}
		temp = temp.Next
		if temp == nil {
			break
		}
	}

	return listRes
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lenHead int // длина списка
	temp := head
	for {
		lenHead++
		temp = temp.Next
		if temp == nil {
			break
		}
	}

	h := true
	res := &ListNode{}
	n = lenHead + 1 - n

	step := 1
	for {
		if n == 1 && step == 1 {
			head = head.Next
		}

		if head == nil {
			break
		}

		step++

		if h {
			res.Val = head.Val
			h = false
		} else {
			res = &ListNode{Val: head.Val, Next: res}
		}

		if step == n {
			head = head.Next
		}
		head = head.Next

		if head == nil {
			break
		}
	}

	// переворачиваем список
	listRes := &ListNode{Val: res.Val}
	res = res.Next
	if res == nil {
		if lenHead == 2 {
			return listRes
		}
		return nil
	}

	for {
		listRes = &ListNode{Val: res.Val, Next: listRes}
		res = res.Next
		if res == nil {
			break
		}
	}

	return listRes
}

// создаем и заполняем список из массива данных
func NewList(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	l := Head(data[0])
	for i := 1; i < len(data); i++ {
		l = AddList(l, data[i])
	}
	return l
}

// создаем головной (первый) узел списка
func Head(d int) *ListNode {
	return &ListNode{
		Val:  d,
		Next: nil,
	}
}

// добавляем к списку один узел
func AddList(l *ListNode, d int) *ListNode {
	return &ListNode{
		Val:  d,
		Next: l,
	}
}

// получаем данные из списка
func getDataList(l *ListNode) []int {
	if l == nil {
		return []int{}
	}
	data := make([]int, 0)
	for {
		data = append(data, l.Val)
		l = l.Next
		if l == nil {
			break
		}
	}
	return data
}
