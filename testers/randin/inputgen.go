package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax int = 1, 10
	mmin, mmax int = 10, 10
)

func CreateInput() string {
	ans := new(strings.Builder)

	N := rand.Intn(nmax-nmin+1) + nmin
	M := rand.Intn(mmax-mmin+1) + mmin
	fmt.Fprintln(ans, N, M)

	for i := 0; i < M; i++ {
		q := rand.Intn(4) + 1
		switch q {
		case 1:
			l := rand.Intn(N) + 1
			r := rand.Intn(N-l+1) + l
			fmt.Fprintln(ans, q, l, r)
		case 2:
			l := rand.Intn(N) + 1
			r := rand.Intn(N-l+1) + l
			x := rand.Intn(2*N+1) - N
			fmt.Fprintln(ans, q, l, r, x)
		case 3:
			i := rand.Intn(N) + 1
			fmt.Fprintln(ans, q, i)
		case 4:
			x := rand.Intn(N) + 1
			fmt.Fprintln(ans, q, x)
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
