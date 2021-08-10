package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CreateInput() string {
	ans := new(strings.Builder)

	N := 15
	fmt.Fprintln(ans, N)
	for i := 0; i < 15; i++ {
		fmt.Fprintln(ans, rand.Intn(10000)+1)
	}
	K := rand.Intn(100) + 1
	fmt.Fprintln(ans, K)

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}
