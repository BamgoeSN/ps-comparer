package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func push(s *[]int, x int) {
	*s = append(*s, x)
}

func get(s *[]int) int {
	if len(*s) == 0 {
		return -1
	}
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	seq := make([]int, n)
	for i := range seq {
		scanner.Scan()
		seq[i], _ = strconv.Atoi(scanner.Text())
	}
	s := new([]int)
	var num = 1
	oper := make([]byte, 0)

	for num <= n {
		if len(*s) == 0 || (*s)[len(*s)-1] != seq[0] {
			push(s, num)
			num++
			oper = append(oper, '+')
		} else {
			get(s)
			seq = seq[1:]
			oper = append(oper, '-')
		}
	}

	solve := true
	for _, v := range seq {
		if v != get(s) {
			solve = false
			break
		}
		oper = append(oper, '-')
	}

	output := make([]byte, 0)
	if solve {
		for _, v := range oper {
			output = append(output, v)
			output = append(output, '\n')
		}
	} else {
		output = append(output, []byte("NO\n")...)
	}
	fmt.Print(string(output))
}
