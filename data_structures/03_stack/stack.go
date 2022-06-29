/*
Стек
	Стек следует принципу Last-In-First-Out (LIFO, “первым на вход — последним на выход”)
    Основные операции со стеком
    - Push — вставка элемента в верхнюю часть стека.
    - Pop — удаление элемента из верхней части стека с возвращением элемента.
    - Peek — просмотр элемента в верхней части стека.
    - isEmpty — проверка пустоты стека.
*/
package main

import (
	"fmt"
)

type Stack struct {
	Val []int
}

func main() {
	s := newStack()
	s.Push(10)
	s.Display()
	fmt.Println(s.Pop())
	fmt.Println("is empty:", s.IsEmpty())
	fmt.Println(s.Peek())
	s.Display()
	s.Push(20)
	s.Push(30)
	s.Display()
	fmt.Println(s.Peek())
	s.Push(40)
	s.Push(50)
	s.Push(60)
	s.Push(70)

	for {
		val, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Printf("%d ", val)
	}
}

//  newStack - создание нового стека
func newStack() *Stack {
	return &Stack{
		Val: []int{},
	}
}

// Push — вставка элемента в верхнюю часть стека
func (s *Stack) Push(val int) {
	s.Val = append(s.Val, val)
}

// Display - вывод всех элементов стека без удаления
func (s *Stack) Display() {
	fmt.Println(s.Val)
}

// Pop — удаление элемента из верхней части стека с возвращением элемента
func (s *Stack) Pop() (int, bool) {
	l := len(s.Val) - 1
	if l == -1 {
		return 0, false
	}
	val := s.Val[l]
	s.Val = s.Val[:l]
	return val, true
}

// Peek — просмотр элемента в верхней части стека
func (s *Stack) Peek() (int, bool) {
	l := len(s.Val) - 1
	if l == -1 {
		return 0, false
	}
	val := s.Val[l]
	return val, true
}

// isEmpty — проверка пустоты стека
func (s *Stack) IsEmpty() bool {
	return len(s.Val) == 0
}
