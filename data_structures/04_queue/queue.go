/*
Очередь
    Основные операции с очередями
    - Enqueue — вставка элемента в конец очереди.
    - Dequeue — удаление элемента из передней части очереди.
    - Top/Peek — возвращает элемент из передней части очереди без удаления.
    - isEmpty — проверка пустоты очереди.
*/

package main

import "fmt"

type Queue struct {
	Val []int
}

func main() {
	q := newQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Display()

	fmt.Println(q.Peek())
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Display()
	fmt.Println(q.IsEmpty())
}

// newQueue- создание новой очереди
func newQueue() *Queue {
	return &Queue{
		Val: []int{},
	}
}

// Enqueue — вставка элемента в конец очереди
func (q *Queue) Enqueue(val int) {
	q.Val = append(q.Val, val)
}

// Display - вывод всех элементов очереди без удаления
func (q *Queue) Display() {
	fmt.Println(q.Val)
}

// Dequeue — удаление элемента из передней части очереди
func (q *Queue) Dequeue() {
	if len(q.Val) == 0 {
		return
	}
	q.Val = q.Val[1:]
}

// Peek — возвращает элемент из передней части очереди без удаления
func (q *Queue) Peek() (int, bool) {
	if len(q.Val) == 0 {
		return 0, false
	}
	val := q.Val[0]
	return val, true
}

// isEmpty — проверка пустоты очереди
func (q *Queue) IsEmpty() bool {
	return len(q.Val) == 0
}
