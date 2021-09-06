package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	smin, smax int = 1, 2
	bmin, bmax int = 0, 100
)

func CreateInput() string {
	ans := new(strings.Builder)

	n := RandRange(smin, smax)
	m := RandRange(smin, smax)
	b := RandRange(bmin, bmax)
	fmt.Fprintln(ans, n, m, b)

	arr := make([][]int, n)
	for r := range arr {
		arr[r] = make([]int, m)
		for c := range arr[r] {
			arr[r][c] = RandRange(0, 25) * 10
			fmt.Fprintf(ans, "%d ", arr[r][c])
		}
		fmt.Fprintln(ans)
	}

	return TrimSpaces(ans.String())
}

func RandRange(l, r int) int {
	return rand.Intn(r-l+1) + l
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}
