package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax int = 0, 100
)

func CreateInput() string {
	ans := new(strings.Builder)

	N := rand.Intn(nmax-nmin+1) + nmin
	fmt.Fprintln(ans, N)

	cnt := 0
	isBroken := make([]bool, 10)
	rate := rand.Float64()
	for i := range isBroken {
		pick := rand.Float64()
		if pick < rate {
			isBroken[i] = true
			cnt++
		}
	}

	fmt.Fprintln(ans, cnt)
	for i, b := range isBroken {
		if b {
			fmt.Fprintf(ans, "%d ", i)
		}
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
