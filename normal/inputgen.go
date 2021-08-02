package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CreateInput() string {
	ans := new(strings.Builder)

	nmax := 10
	pathmax := 10
	costmax := 10

	N := rand.Intn(nmax-1) + 2
	fmt.Fprintln(ans, N)

	for i := 0; i < N-1; i++ {
		fmt.Fprintf(ans, "%d ", rand.Intn(pathmax)+1)
	}
	fmt.Fprintln(ans)
	for i := 0; i < N; i++ {
		fmt.Fprintf(ans, "%d ", rand.Intn(costmax)+1)
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
