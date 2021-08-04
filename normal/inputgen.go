package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax int = 1, 5
	mmin, mmax int = 1, 10
)

func CreateInput() string {
	ans := new(strings.Builder)

	n := rand.Intn(nmax-nmin+1) + nmin
	m := rand.Intn(mmax-mmin+1) + mmin
	v := rand.Intn(n) + 1
	fmt.Fprintln(ans, n, m, v)

	for i := 0; i < m; i++ {
		s := rand.Intn(n)
		e := rand.Intn(n)
		// if e >= s {
		// 	e++
		// }
		fmt.Fprintln(ans, s+1, e+1)
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
