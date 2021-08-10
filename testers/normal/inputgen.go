package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmax, nmin int = 100, 1
)

func CreateInput() string {
	ans := new(strings.Builder)

	M := rand.Intn(nmax-nmin+1) + nmin
	N := rand.Intn(M-nmin+1) + nmin
	fmt.Fprintf(ans, "%d\n%d\n", M, N)

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}
