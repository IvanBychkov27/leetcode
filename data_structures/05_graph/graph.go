/*
Граф
    Ключевые термины
    - Размер — количество ребер в графе.
    - Порядок — количество вершин в графе.
    - Смежность — случай, когда два узла соединены одним и тем же ребром.
    - Петля — вершина, соединенная ребром сама с собой.
    - Изолированная вершина — вершина, которая не связана с другими вершинами.

    Графы делятся на два типа. Они различаются главным образом по направлениям пути между двумя вершинами.
    - Ориентированные графы: все ребра имеют направления, указывающие начальную и конечную точки (вершины).
    - Неориентированные графы: ребра не имеют направлений, которые позволяют обходам происходить с любого направления.

    Распространенные алгоритмы обхода графов
    - Поиск в ширину (BFS) — метод поиска кратчайшего пути в графе, основанный на вершинах.
    - Поиск в глубину (DFS) — метод, основанный на ребрах.

    Основные операции с графами
    - Add vertex: добавить вершину в граф.
    - Add edge: добавить ребро между двумя вершинами.
    - Display: отобразить вершину.
    - Total cost of traversal: найти общую стоимость пути обхода.
*/

package main

import (
	"fmt"
	"sort"
)

// Общая инфо о графе
type GraphInfo struct {
	Head        *Vertex         // указатель на головную вершину (голова)
	Tail        *Vertex         // указатель на последнюю вершину (хвост)
	CountVertex int             // кол-во вершин в графе
	CountAdges  int             // кол-во ребер в графе
	AllPointers map[int]*Vertex // указатели на все вершины графа map[id]*Vertex
}

// Ребро графа
type Adge struct {
	ID          int     // имя узла к которому направлено ребро
	Rate        int     // вес или стоимость ребра
	Pointer     *Vertex // указатель на узел к которому направлено ребро
	Orientation bool    // ориентация ребра: false/true - входящее/исходящее ребро
}

// Вершина графа
type Vertex struct {
	ID    int        // название вершины графа
	Val   int        // значение вершины графа
	Adges []Adge     // рёбра - это с какими вершинами соединена данная вершина
	Info  *GraphInfo // общая информация о графе

}

func main() {
	g := NewGraph(0)
	g = g.AddVertex(1)
	g = g.AddVertex(2)
	g = g.AddVertex(3)
	g = g.AddVertex(4)
	g = g.AddVertex(5)
	g = g.AddVertex(6)
	g = g.AddVertex(7)

	g.AddEdge(0, 1, 100)
	g.AddEdge(1, 3, 300)
	g.AddEdge(3, 7, 300)
	g.AddEdge(1, 4, 400)
	g.AddEdge(4, 7, 300)
	g.AddEdge(4, 7, 300)

	g.AddEdge(0, 2, 200)
	g.AddEdge(2, 5, 100)
	g.AddEdge(5, 7, 400)
	g.AddEdge(2, 6, 300)
	g.AddEdge(6, 7, 100)

	g.AddEdge(5, 6, 100)

	g.DisplayGraph()
	fmt.Println()

	g.DisplayInfoGraph()
	fmt.Println()

	min, route := g.MinRoute(g.Info.Head, []int{})
	fmt.Println("min:", min)
	fmt.Println("ids:", route)
}

// создание нового графа
func NewGraph(val int) *Vertex {
	v := &Vertex{
		ID:    0,
		Val:   val,
		Adges: []Adge{},
	}

	graphInfo := &GraphInfo{
		Head:        v,
		Tail:        v,
		CountVertex: 1,
		CountAdges:  0,
		AllPointers: map[int]*Vertex{v.ID: v},
	}

	v.Info = graphInfo

	return v
}

// AddVertex - добавить вершину в граф
func (v *Vertex) AddVertex(val int) *Vertex {
	v = v.Info.Tail
	n := &Vertex{
		ID:    v.ID + 1,
		Val:   val,
		Adges: []Adge{},
		Info:  v.Info,
	}
	n.Info.Tail = n
	n.Info.CountVertex++
	n.Info.AllPointers[n.ID] = n

	return n
}

// Вывод информации графа по всем вершинам
func (v *Vertex) DisplayGraph() {
	k := make([]int, 0, len(v.Info.AllPointers))
	for _, n := range v.Info.AllPointers {
		k = append(k, n.ID)
	}

	sort.Ints(k)

	for _, id := range k {
		n := v.getVertex(id)
		fmt.Printf("ID:%2d  Val:%3d  Adges:%2d ", n.ID, n.Val, len(n.Adges))
		for _, adge := range n.Adges {
			fmt.Printf("[id %d rate %d %t] ", adge.ID, adge.Rate, adge.Orientation)
		}
		fmt.Printf("\n")
	}
}

// Вывод информации по ID вершины
func (v *Vertex) DisplayVertex(id int) {
	n, ok := v.Info.AllPointers[id]
	if !ok {
		fmt.Printf("ID:%2d no found \n", id)
		return
	}
	fmt.Printf("ID:%2d  Val:%3d  Adges:%2d\n", n.ID, n.Val, len(n.Adges))
}

// DisplayInfoGraph - выводит общую информацию графа
func (v *Vertex) DisplayInfoGraph() {
	fmt.Printf("Head ID: %d,  Tail ID: %d, CountVertex: %d, CountAdges: %d\n",
		v.Info.Head.ID, v.Info.Tail.ID, v.Info.CountVertex, v.Info.CountAdges)
}

// AddEdge - добавляет/обновляет ребро между двумя вершинами idN (начало), idK (конец), с весом rate
func (v *Vertex) AddEdge(idN, idK, rate int) {
	n := v.getVertex(idN)
	k := v.getVertex(idK)

	for i, a := range n.Adges {
		if a.ID == idK {
			n.Adges[i].Rate = rate

			for j, b := range k.Adges {
				if b.ID == idN {
					k.Adges[j].Rate = rate
				}
			}

			return
		}
	}

	n.Adges = append(n.Adges, Adge{
		ID:          idK,
		Rate:        rate,
		Pointer:     v.getVertex(idK),
		Orientation: true,
	})

	k.Adges = append(k.Adges, Adge{
		ID:          idN,
		Rate:        rate,
		Pointer:     n,
		Orientation: false,
	})

	v.Info.CountAdges++
}

// получить указатель на вершину с индексом id
func (v *Vertex) getVertex(id int) *Vertex {
	return v.Info.AllPointers[id]
}

// TotalValueGraph - найти общую вес всех вершин графа
func (v *Vertex) TotalValueGraph() int {
	totalRate := 0
	for _, n := range v.Info.AllPointers {
		totalRate += n.Val
	}
	return totalRate
}

// DeleteAdge - удалить ребро графа
func (v *Vertex) DeleteAdge(idN, idK int) {
	n := v.getVertex(idN)
	if n == nil {
		return
	}
	n.Adges = delAdge(idK, n.Adges)

	k := v.getVertex(idK)
	if k == nil {
		return
	}
	k.Adges = delAdge(idN, k.Adges)
	v.Info.CountAdges--
}

func delAdge(id int, Adges []Adge) []Adge {
	for i, a := range Adges {
		if a.ID == id {
			resAdges := Adges[:i]
			if i+1 <= len(Adges) {
				resAdges = append(resAdges, Adges[i+1:]...)
			}
			return resAdges
		}
	}
	return nil
}

// DeleteVertex - удалить вершину из графа и все связанные с ним ребра
func (v *Vertex) DeleteVertex(id int) {
	n := v.getVertex(id)
	if n == nil {
		return
	}
	for _, a := range n.Adges {
		v.DeleteAdge(id, a.ID)
	}
	delete(v.Info.AllPointers, id)
	v.Info.CountVertex--

	if v.Info.CountVertex == 0 {
		v.Info.Head = nil
		v.Info.Tail = nil
		return
	}

	if id == v.Info.Head.ID {
		headID, _ := v.getNeighborsID(id)
		v.Info.Head = v.getVertex(headID)
	}

	if id == v.Info.Tail.ID {
		_, tailID := v.getNeighborsID(id)
		v.Info.Tail = v.getVertex(tailID)
	}
}

// взять соседние ID
func (v *Vertex) getNeighborsID(id int) (headID, tailID int) {
	headID = len(v.Info.AllPointers)
	for _, n := range v.Info.AllPointers {
		if n.ID < headID {
			headID = n.ID
		}
		if n.ID > tailID {
			tailID = n.ID
		}
	}
	return
}

// MinRoute - поиск минимального пути (маршрута) от вершины Head до Tail в ориентированном графе
func (v *Vertex) MinRoute(n *Vertex, route []int) (int, []int) {
	if n.ID == v.Info.Tail.ID {
		return 0, nil
	}
	if len(n.Adges) == 0 {
		return 0, nil
	}
	id := []int{n.Adges[0].ID}
	min := 1000000000
	for _, a := range n.Adges {
		if !a.Orientation {
			continue
		}
		rate, r := v.MinRoute(a.Pointer, route)

		if min > a.Rate+rate {
			id = []int{a.ID}
			id = append(id, r...)
			min = a.Rate + rate
		}
	}
	route = append(route, id...)

	return min, route
}
