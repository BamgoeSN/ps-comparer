package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//------
// Main
//------

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	N, M, V := nextInt(), nextInt(), nextInt()-1
	graph := make([][]int, N)
	for i := 0; i < M; i++ {
		s, e := nextInt()-1, nextInt()-1
		graph[s] = append(graph[s], e)
		graph[e] = append(graph[e], s)
	}
	for k := range graph {
		sort.Slice(graph[k], func(i, j int) bool { return graph[k][i] < graph[k][j] })
	}

	// DFS
	visit := make([]bool, N)
	visit[V] = true
	list := []int{V}
	DFS(graph, V, visit, &list)
	for _, v := range list {
		fmt.Fprintf(wr, "%d ", v+1)
	}
	fmt.Fprintln(wr)

	// BFS
	q := NewQueue(1)
	q.Push(V)
	for i := range visit {
		visit[i] = false
	}
	visit[V] = true
	fmt.Fprintf(wr, "%d ", V+1)
	for q.Len() != 0 {
		v := q.Pop()
		for _, e := range graph[v] {
			if visit[e] {
				continue
			}
			visit[e] = true
			fmt.Fprintf(wr, "%d ", e+1)
			q.Push(e)
		}
	}
}

func DFS(graph [][]int, curr int, visit []bool, list *[]int) {
	for _, e := range graph[curr] {
		if visit[e] {
			continue
		}
		visit[e] = true
		*list = append(*list, e)
		DFS(graph, e, visit, list)
	}
}

//-------
// Queue
//-------

type Queue struct {
	ints        []int
	front, rear int
	cap         int
}

func NewQueue(cap int) *Queue {
	if cap <= 1 {
		cap = 2
	}
	q := new(Queue)
	q.ints = make([]int, cap)
	q.front, q.rear = 0, 0
	q.cap = cap
	return q
}

func (q *Queue) extend(newCap int) {
	newLen := q.Len()
	newArr := make([]int, newCap)
	if q.rear >= q.front {
		copy(newArr, q.ints[q.front:q.rear])
	} else {
		lenFront := len(q.ints) - q.front
		copy(newArr, q.ints[q.front:])
		copy(newArr[lenFront:], q.ints[:q.rear])
	}
	q.ints = newArr
	q.cap = newCap
	q.front, q.rear = 0, newLen
}

func (q *Queue) Push(n int) {
	if q.Len() >= q.cap-1 {
		q.extend(q.cap*3/2 + 1)
	}
	q.ints[q.rear] = n
	q.rear++
	if q.rear >= q.cap {
		q.rear -= q.cap
	}
}

func (q *Queue) Pop() int {
	temp := q.ints[q.front]
	q.front++
	if q.front >= q.cap {
		q.front -= q.cap
	}
	return temp
}

func (q *Queue) Len() int {
	length := q.rear - q.front + q.cap
	if length >= q.cap {
		length -= q.cap
	}
	return length
}

//---------
// Fast IO
//---------

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	text := sc.Text()
	v, _ := strconv.Atoi(text)
	return v
}
