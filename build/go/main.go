package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var input [10]int64
	for i := range input {
		fmt.Fscan(r, &(input[i]))
	}

	board := make(map[int64]bool)
	for _, v := range input {
		board[v%42] = true
	}

	fmt.Println(len(board))
}
