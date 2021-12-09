/* https://leetcode.com/problems/merge-two-sorted-lists/
Вам даны заголовки двух отсортированных связанных списков list1 и list2.
Объедините два списка в один отсортированный список.
Список должен быть составлен путем соединения узлов первых двух списков.
Верните заголовок объединенного связанного списка.

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправленных онлайн-сообщений для объединения двух отсортированных списков.
Использование памяти: 2,8 МБ, менее 19,18 % отправленных онлайн-сообщений для объединения двух отсортированных списков.

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//d1, d2 := []int{1, 2, 4}, []int{1, 3, 4}
	d1, d2 := []int{}, []int{}
	//d1, d2 := []int{}, []int{0}

	l1 := NewList(d1)
	l2 := NewList(d2)
	fmt.Println("l1 =", getDataList(l1))
	fmt.Println("l2 =", getDataList(l2))

	l := mergeTwoLists(l1, l2)
	fmt.Println("l  =", getDataList(l))

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var l *ListNode
	for {
		if list1 == nil && list2 == nil {
			break
		}

		if list1 == nil && list2 != nil {
			l = &ListNode{Val: list2.Val, Next: l}
			list2 = list2.Next
		}

		if list2 == nil && list1 != nil {
			l = &ListNode{Val: list1.Val, Next: l}
			list1 = list1.Next
		}

		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				l = &ListNode{Val: list1.Val, Next: l}
				list1 = list1.Next
			} else {
				l = &ListNode{Val: list2.Val, Next: l}
				list2 = list2.Next
			}
		}
	}

	if l == nil {
		return nil
	}
	// переворачиваем список
	listRes := &ListNode{Val: l.Val}
	l = l.Next
	if l == nil {
		return listRes
	}

	for {
		listRes = &ListNode{Val: l.Val, Next: listRes}
		l = l.Next
		if l == nil {
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
	l := Head(data[len(data)-1])
	for i := len(data) - 2; i > -1; i-- {
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
