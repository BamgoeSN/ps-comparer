package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//---------------
// Board tilting
//---------------
var board [][]byte

type Stat struct {
	depth          int
	rr, rc, br, bc int
}

const (
	Up = iota
	Down
	Left
	Right
)

var dir = []int{Up, Down, Left, Right}
var dirArr = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func move(tr, tc, or, oc *int, dir int) (isDropped bool) {
	if *tr == -1 && *tc == -1 {
		return true
	}

	isDropped = false
	rp, cp := dirArr[dir][0], dirArr[dir][1]

	for (*tr+rp != *or || *tc+cp != *oc) && (board[*tr+rp][*tc+cp] != '#') {
		*tr += rp
		*tc += cp
		if board[*tr][*tc] == 'O' {
			isDropped = true
			*tr, *tc = -1, -1
			break
		}
	}
	return
}

func tilt(curr Stat, dir int) (Stat, int) {
	var isDropped bool
	res := 0
	curr.depth++

	isDropped = move(&(curr.rr), &(curr.rc), &(curr.br), &(curr.bc), dir)
	if isDropped {
		res = 1
	}

	isDropped = move(&(curr.br), &(curr.bc), &(curr.rr), &(curr.rc), dir)
	if isDropped {
		res = -1
	}

	isDropped = move(&(curr.rr), &(curr.rc), &(curr.br), &(curr.bc), dir)
	if isDropped && res != -1 {
		res = 1
	}

	return curr, res
}

//-------
// Queue
//-------

type Queue struct {
	item        []*Stat
	front, rear int
	cap         int
}

func NewQueue(cap int) *Queue {
	q := new(Queue)
	q.item = make([]*Stat, cap)
	q.front, q.rear = 0, 0
	q.cap = cap
	return q
}

func (q *Queue) extend(newCap int) {
	newLen := q.Len()
	newArr := make([]*Stat, newCap)
	if q.rear >= q.front {
		copy(newArr, q.item[q.front:q.rear])
	} else {
		lenFront := len(q.item) - q.front
		copy(newArr, q.item[q.front:])
		copy(newArr[lenFront:], q.item[:q.rear])
	}
	q.item = newArr
	q.cap = newCap
	q.front, q.rear = 0, newLen
}

func (q *Queue) Push(n *Stat) {
	// If the queue is about to overflow, reallocate data
	if q.Len() >= q.cap-1 {
		q.extend(q.cap*3/2 + 1)
	}

	q.item[q.rear] = n
	q.rear++
	q.rear %= q.cap
}

func (q *Queue) Pop() *Stat {
	temp := q.item[q.front]
	q.front++
	q.front %= q.cap
	return temp
}

func (q *Queue) Len() int {
	length := (q.rear - q.front) % q.cap
	if length < 0 {
		length += q.cap
	}
	return length
}

//------
// Main
//------

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	// Inputs
	size := readInts()
	board = make([][]byte, size[0])

	for i := range board {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		board[i] = []byte(text)
	}

	// Find positions of red and blue
	start := Stat{0, -1, -1, -1, -1}
	for i, row := range board {
		for j, elm := range row {
			if elm == 'R' {
				board[i][j] = '.'
				start.rr, start.rc = i, j
			} else if elm == 'B' {
				board[i][j] = '.'
				start.br, start.bc = i, j
			}
		}
	}

	// BFS
	visited := make([][][][]bool, size[0])
	for i := range visited {
		visited[i] = make([][][]bool, size[1])
		for j := range visited[i] {
			visited[i][j] = make([][]bool, size[0])
			for k := range visited[i][j] {
				visited[i][j][k] = make([]bool, size[1])
			}
		}
	}

	q := NewQueue(2)
	q.Push(&start)

	var buf Stat
	var found bool = false
	for q.Len() != 0 {
		curr := q.Pop()
		if visited[curr.rr][curr.rc][curr.br][curr.bc] {
			continue
		}
		visited[curr.rr][curr.rc][curr.br][curr.bc] = true

		if curr.depth >= 10 {
			continue
		}

		for _, d := range dir {
			next, res := tilt(*curr, d)
			switch res {
			case 0:
				q.Push(&next)
			case 1:
				buf = next
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	if found {
		fmt.Fprintln(writer, buf.depth)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func readInts() []int {
	text, _ := reader.ReadString('\n')
	fld := strings.Fields(text)
	res := make([]int, len(fld))
	for i, v := range fld {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}
