package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

var obj = []byte("RBO")

func CreateInput() string {
	ans := new(strings.Builder)

	minsize := 4
	maxsize := 10

	R := rand.Intn(maxsize-minsize+1) + minsize
	C := rand.Intn(maxsize-minsize+1) + minsize
	fmt.Fprintln(ans, R, C)

	board := make([][]byte, R)
	for i := range board {
		board[i] = make([]byte, C)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// Create surrounding walls
	for r := 0; r < R; r++ {
		board[r][0], board[r][C-1] = '#', '#'
	}
	for c := 0; c < C; c++ {
		board[0][c], board[R-1][c] = '#', '#'
	}

	// Place required objects
	for _, b := range obj {
		r, c := 0, 0
		for board[r][c] != '.' {
			r = rand.Intn(R-2) + 1
			c = rand.Intn(C-2) + 1
			// fmt.Println(R, C, r, c)
		}
		board[r][c] = b
	}

	// Place walls
	rate := rand.Float64()
	rate = math.Pow(rate, 5)
	for r := 1; r < R-1; r++ {
		for c := 1; c < C-1; c++ {
			if board[r][c] != '.' {
				continue
			}
			pick := rand.Float64()
			if pick < rate {
				board[r][c] = '#'
			}
		}
	}

	for _, v := range board {
		fmt.Fprintln(ans, string(v))
	}

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}
