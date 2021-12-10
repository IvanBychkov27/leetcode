/* https://leetcode.com/problems/merge-k-sorted-lists/

Вам предоставляется массив из k связанных списков, каждый связанный список отсортирован в порядке возрастания.
Объедините все связанные списки в один отсортированный связанный список и верните его.

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Выполнено:
Время выполнения: 264 мс, быстрее, чем 18,02 % отправленных онлайн-заявок на слияние k отсортированных списков.
Использование памяти: 6,1 МБ, менее 45,05 % отправленных онлайн-заявок на слияние k отсортированных списков.

*/

package main

import "fmt"

func main() {
	d1, d2, d3 := []int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}
	//d1, d2, d3 := []int{}, []int{}, []int{}

	l1 := NewList(d1)
	l2 := NewList(d2)
	l3 := NewList(d3)

	fmt.Println("l1 =", getDataList(l1))
	fmt.Println("l2 =", getDataList(l2))
	fmt.Println("l3 =", getDataList(l3))

	l := mergeKLists([]*ListNode{l1, l2, l3})

	fmt.Println("l  =", getDataList(l))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var l *ListNode
	lenLists := len(lists)

	for {
		min := 100000
		idx := 0
		count := 0
		for i, list := range lists {
			if list == nil {
				count++
				continue
			}
			if list.Val < min {
				min = list.Val
				idx = i
			}
		}

		if count == lenLists {
			break
		}

		l = &ListNode{Val: lists[idx].Val, Next: l}
		lists[idx] = lists[idx].Next

	}

	return listFlip(l)
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
