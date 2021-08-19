package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax = 10, 10
	dir        = "UDLR"
	dr         = []int{-1, 1, 0, 0}
	dc         = []int{0, 0, -1, 1}
)

func CreateInput() string {
	ans := new(strings.Builder)

	R := rand.Intn(nmax-nmin+1) + nmin
	C := rand.Intn(nmax-nmin+1) + nmin
	if R == 1 && C == 1 {
		pick := rand.Intn(2)
		if pick == 1 {
			R = 2
		} else {
			C = 2
		}
	}
	fmt.Fprintln(ans, R, C)

	arr := make([][]byte, R)
	for r := range arr {
		arr[r] = make([]byte, C)
		for c := range arr[r] {
			v := -1
			for v < 0 || (r+dr[v] < 0 || r+dr[v] >= R || c+dc[v] < 0 || c+dc[v] >= C) {
				v = rand.Intn(4)
			}
			arr[r][c] = dir[v]
		}
		fmt.Fprintln(ans, string(arr[r]))
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
