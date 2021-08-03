package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var n int = 1

func CreateInput() string {
	ans := new(strings.Builder)

	nmax := 42
	for i := 0; i < 10; i++ {
		fmt.Fprintln(ans, rand.Intn(nmax)+1)
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
