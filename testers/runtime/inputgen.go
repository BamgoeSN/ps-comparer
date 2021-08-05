package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var ()

func CreateInput() string {
	ans := new(strings.Builder)

	N := 1000000
	fmt.Fprintln(ans, N)
	for i := 0; i < N; i++ {
		fmt.Fprintf(ans, "%d ", rand.Intn(100000)+1)
	}
	fmt.Fprintln(ans)

	M := 100000
	fmt.Fprintln(ans, M)
	for i := 0; i < M; i++ {
		e := rand.Intn(N-2) + 2
		s := rand.Intn(e-1) + 1
		fmt.Fprintln(ans, s, e)
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
