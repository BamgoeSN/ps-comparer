package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//-------
// Stack
//-------

type Stack []int64

func NewStack(cap int) *Stack {
	s := make(Stack, 0, cap)
	return &s
}

func (s *Stack) Push(n int64) {
	*s = append(*s, n)
}

func (s *Stack) Pop() int64 {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *Stack) Get() int64 {
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

//------
// Main
//------

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	N := nextInt()
	s := NewStack(2)

	for i := 0; i < N; i++ {
		sc.Scan()
		q := sc.Text()
		switch q {
		case "push":
			s.Push(nextInt64())
		case "pop":
			if s.Len() == 0 {
				fmt.Fprint(wr, "-1\n")
			} else {
				fmt.Fprintf(wr, "%d\n", s.Pop())
			}
		case "size":
			fmt.Fprintf(wr, "%d\n", s.Len())
		case "empty":
			if s.Len() == 0 {
				fmt.Fprint(wr, "1\n")
			} else {
				fmt.Fprint(wr, "0\n")
			}
		case "top":
			if s.Len() == 0 {
				fmt.Fprint(wr, "-1\n")
			} else {
				fmt.Fprintf(wr, "%d\n", s.Get())
			}
		}
	}
}

//---------
// Fast IO
//---------

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() (res int) {
	sc.Scan()
	text := sc.Text()
	v, _ := strconv.Atoi(text)
	return v
}

func nextInt64() (res int64) {
	sc.Scan()
	text := sc.Text()
	v, _ := strconv.ParseInt(text, 10, 64)
	return v
}
