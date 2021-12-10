/* https://leetcode.com/problems/swap-nodes-in-pairs/
Учитывая связанный список, поменяйте местами каждые два соседних узла и верните его заголовок.
Вы должны решить проблему, не изменяя значения в узлах списка (т.е. могут быть изменены только сами узлы).

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправлений в режиме онлайн для узлов подкачки в парах.
Использование памяти: 2,1 МБ, менее 33,99 % отправленных онлайн-сообщений для узлов подкачки в парах.

*/

package main

import (
	"fmt"
)

func main() {
	//d := []int{1, 2, 3, 4, 5, 6, 7}
	d := []int{1, 2, 3, 4, 5, 6}
	//d := []int{1, 2}
	//d := []int{1}
	//d := []int{}
	l := NewList(d)
	fmt.Println("l   =", getDataList(l))

	res := swapPairs(l)
	fmt.Println("res =", getDataList(res))

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	lenRes := lenList(head)
	if lenRes == 1 {
		return head
	}

	if lenRes%2 != 0 {
		lenRes--
	}
	res := swap(head, 0)
	for i := 2; i < lenRes; i += 2 {
		swap(res, i)
	}

	return res
}

func swap(head *ListNode, n int) *ListNode {
	if n == 0 {
		a := head
		b := a.Next
		a.Next, b.Next = b.Next, a
		return b
	}

	for i := 0; i < n-1; i++ {
		head = head.Next
	}

	a := head.Next
	b := a.Next
	head.Next = b
	a.Next, b.Next = b.Next, a
	return b
}

// длина списка
func lenList(l *ListNode) int {
	count := 0
	for {
		if l == nil {
			return count
		}
		count++
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// переворачиваем список
func listFlip(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	var listRes *ListNode
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
