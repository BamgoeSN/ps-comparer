package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	vmin, vmax int = -100, 100
	nmin, nmax int = 1, 100
)

func CreateInput() string {
	ans := new(strings.Builder)

	n := RandRange(nmin, nmax)
	fmt.Fprintln(ans, n)

	for i := 0; i < n; i++ {
		v := RandRange(vmin, vmax)
		fmt.Fprintf(ans, "%d ", v)
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

func RandRange(l, r int) int       { return rand.Intn(r-l+1) + l }
func RandRange64(l, r int64) int64 { return rand.Int63n(r-l+1) + l }
