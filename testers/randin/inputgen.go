package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmax, nmin int = 8, 1
)

func CreateInput() string {
	ans := new(strings.Builder)

	R, C := rand.Intn(nmax-nmin+1)+nmin, rand.Intn(nmax-nmin+1)+nmin
	fmt.Fprintln(ans, R, C)

	rate := rand.Float64()
	arr := make([][]byte, R)
	for r := range arr {
		arr[r] = make([]byte, C)
		for c := range arr[r] {
			pick := rand.Float64()
			if pick < rate {
				arr[r][c] = '1'
			} else {
				arr[r][c] = '0'
			}
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
