/*
Runtime: 4 ms, faster than 97.86% of Go online submissions for Add Two Numbers.
Memory Usage: 5.1 MB, less than 16.71% of Go online submissions for Add Two Numbers.
*/
// https://leetcode.com/problems/add-two-numbers/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//d1 := []int{0}
	//d2 := []int{0}

	d1 := []int{2, 4, 3}
	d2 := []int{5, 6, 4}

	//d1 := []int{9, 9, 9, 9, 9, 9, 9}
	//d2 := []int{9, 9, 9, 9}

	//d1 := []int{9, 9, 9, 9}
	//d2 := []int{9, 9, 9, 9, 9, 9, 9}

	l1 := NewList(d1)
	l2 := NewList(d2)

	fmt.Println(getDataList(l1))
	fmt.Println(getDataList(l2))

	l := addTwoNumbers(l1, l2)
	fmt.Println(getDataList(l))
}

// узел списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// складываем числа из двух списков (каждый список представляет число - в узле списка одна цифра числа)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var up, v1, v2 int
	head := true
	l := &ListNode{}
	for {
		v1, v2 = 0, 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		data := v1 + v2 + up

		up = 0
		s := strconv.Itoa(data)
		if len(s) == 2 {
			up, _ = strconv.Atoi(string(s[0]))
			data, _ = strconv.Atoi(string(s[1]))
		}

		if head {
			l.Val = data
			head = false
		} else {
			l = &ListNode{Val: data, Next: l}
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if up > 0 {
				l = &ListNode{Val: up, Next: l}
			}
			break
		}

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
