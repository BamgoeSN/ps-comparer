package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CreateInput() string {
	ans := new(strings.Builder)

	N := 10
	K := RandRange(1, N)
	fmt.Fprintln(ans, N, K)

	for i := 0; i < N; i++ {
		fmt.Fprintln(ans, RandRange(0, 30))
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
