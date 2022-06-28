/*
Связанный список (Linked List)
    - Односвязный. Обход элементов может выполняться только в прямом направлении.
    - Двусвязный. Обход элементов может выполняться как в прямом, так и в обратном направлениях. Узлы включают дополнительный указатель, известный как prev, указывающий на предыдущий узел.
    - Круговые связанные. Это связанные списки, в которых предыдущий (prev) указатель “головы” указывает на “хвост”, а следующий указатель “хвоста” указывает на “голову”.

    Основные операции со связанными списками
    - Insertion — добавление узла в список. Это может быть сделано на основе требуемого местоположения, такого как голова, хвост или где-то посередине.
    - Delete — удаление узла в начале списка или на основе заданного ключа.
    - Display — отображение полного списка.
    - Search — поиск узла в данном связанном списке.
    - Update — обновление значения узла в заданном ключе.
*/

package main

import "fmt"

type ListInfo struct {
	Head *ListNode // указатель на головной узел (голова)
	Tail *ListNode // указатель на последний узел (хвост)
	Len  int       // длина списка
}

// узел списка
type ListNode struct {
	ID   int       // индекс узла
	Val  int       // данные узла
	Info *ListInfo // информация по списку
	//Head *ListNode // указатель на головной узел (голова)
	Next *ListNode // указатель на следующий узел
	Prev *ListNode // указатель на предыдущий узел
	//Tail *ListNode // указатель на последний узел (хвост)
	//Len  int       // длина списка
}

func main() {
	l := Head(10)
	l = l.AddList(20)
	l = l.AddList(30)
	l = l.AddList(40)
	l = l.AddList(50)
	l = l.AddList(60)
	l.DisplayList()
	l.Reindexing()

	var ok bool
	l, ok = l.Search(50)
	if ok {
		l.DisplayNode()
	}

	l, ok = l.Update(l.ID, 500)
	if ok {
		l.DisplayNode()
	}

	l, ok = l.Insertion(0, 0)
	l, ok = l.Insertion(2, 22)
	l, ok = l.Insertion(5, 55)
	l, ok = l.Insertion(10, 99)
	if ok {
		l.DisplayList()
	}

	l, ok = l.Delete(7)
	if ok {
		l.DisplayList()
	}
}

// создаем головной (первый) узел списка
func Head(val int) *ListNode {
	l := &ListNode{
		ID:   0,
		Val:  val,
		Next: nil,
		Prev: nil,
	}
	info := &ListInfo{
		Head: l,
		Tail: l,
		Len:  1,
	}
	l.Info = info
	return l
}

// в конец списка добавляем один узел
func (l *ListNode) AddList(val int) *ListNode {
	l = l.Info.Tail // перемещаемся в хвост списка
	n := &ListNode{
		ID:   l.ID + 1,
		Val:  val,
		Info: l.Info,
		Next: nil,
		Prev: l,
	}
	l.Next = n
	l.Info.Tail = n
	l.Info.Len++
	return n
}

// Insertion — добавление узла в список.
// Это может быть сделано на основе требуемого местоположения,
// такого как голова, хвост или где-то посередине
func (l *ListNode) Insertion(id, val int) (*ListNode, bool) {
	if id < 0 {
		return l, false
	}
	l = l.Info.Head
	if id == 0 {
		n := &ListNode{
			ID:   0,
			Val:  val,
			Info: l.Info,
			Next: l,
			Prev: nil,
		}
		n.Info.Head = n
		n.Reindexing()
		return n, true
	}

	if id >= l.LenList() {
		return l.AddList(val), true
	}

	var prev *ListNode
	for {
		if id == l.ID {
			n := &ListNode{
				ID:   id,
				Val:  val,
				Info: l.Info,
				Next: l,
				Prev: prev,
			}
			prev.Next = n
			n.Reindexing()
			return n, true
		}

		if l.Next == nil {
			return l, false
		}
		prev = l
		l = l.Next
	}
}

// Search — поиск узла в списке
func (l *ListNode) Search(val int) (*ListNode, bool) {
	l = l.Info.Head
	for {
		if l.Val == val {
			return l, true
		}
		if l.Next == nil {
			return l.Info.Head, false
		}
		l = l.Next
	}
}

// Update — обновление значения узла в заданном ключе
func (l *ListNode) Update(id, val int) (*ListNode, bool) {
	l = l.Info.Head
	for {
		if l.ID == id {
			l.Val = val
			return l, true
		}
		if l.Next == nil {
			return l.Info.Head, false
		}
		l = l.Next
	}
}

func (l *ListNode) DeleteList() {
	l = l.Info.Head
	l.ID = 0
	l.Val = 0
	l.Next = nil
	l.Prev = nil
	l.Info.Head = nil
	l.Info.Tail = nil
	l.Info.Len = 0
}

// Delete — удаление узла в начале списка или на основе заданного ключа
func (l *ListNode) Delete(id int) (*ListNode, bool) {
	l = l.Info.Head
	if id == 0 || l.LenList()-1 == 0 { // удаление головного узла
		if l.Next == nil {
			l.ID = 0
			l.Val = 0
			l.Info.Head = nil
			l.Info.Tail = nil
			l.Info.Len = 0
			return l, true
		}
		l = l.Next
		l.Prev = nil
		l.Info.Head = l
		l.Info.Len--
		return l, true
	}

	if id >= l.LenList()-1 {
		l = l.Info.Tail
		l = l.Prev
		l.Next = nil
		l.Info.Tail = l
		l.Info.Len--
		return l, true
	}

	var prev *ListNode
	for {
		if l.ID == id {
			l = l.Next
			l.Prev = prev
			prev.Next = l
			l.Info.Len--
			return l, true
		}
		if l.Next == nil {
			return l, false
		}
		prev = l
		l = l.Next
	}
}

// распечатать один узел из списка
func (l *ListNode) DisplayNode() {
	if l == nil {
		return
	}
	fmt.Printf("id: %d  val: %d\n", l.ID, l.Val)
}

// Display — отображение полного списка начиная с головы
func (l *ListNode) DisplayList() {
	if l == nil {
		return
	}
	l = l.Info.Head
	if l == nil {
		fmt.Printf("list nil")
		return
	}
	for {
		fmt.Printf("[%d:%3d] ", l.ID, l.Val)
		if l.Next == nil {
			fmt.Printf(" (len %d)\n", l.LenList())
			return
		}
		l = l.Next
	}
}

// длина списка
func (l *ListNode) LenList() int {
	return l.Info.Len
}

// переиндексация списка
func (l *ListNode) Reindexing() {
	if l == nil {
		return
	}
	l = l.Info.Head
	id := 0
	for {
		l.ID = id
		if l.Next == nil {
			l.Info.Len = id + 1
			return
		}
		l = l.Next
		id++
	}
}
