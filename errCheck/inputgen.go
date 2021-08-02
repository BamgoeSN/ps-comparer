package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CreateInput() string {
	ans := new(strings.Builder)

	TC := rand.Intn(10)
	nmax := 10
	wmax := int64(10)

	for TC != 0 {
		TC--

		N, M := rand.Intn(nmax-1)+2, rand.Intn(nmax)+1
		fmt.Fprintln(ans, N, M)

		list := make(map[[2]int]bool)

		for i := 0; i < M; i++ {
			pick := rand.Intn(2)
			if pick == 1 {
				fmt.Fprint(ans, "! ")
				a, b, w := rand.Intn(nmax)+1, rand.Intn(nmax)+1, rand.Int63n(wmax+1)
				for {
					if _, e := list[[2]int{a, b}]; !e {
						break
					}
					a, b = rand.Intn(nmax)+1, rand.Intn(nmax)+1
				}
				list[[2]int{a, b}] = true
				fmt.Fprintln(ans, a, b, w)
			} else {
				fmt.Fprint(ans, "? ")
				a, b := rand.Intn(nmax)+1, rand.Intn(nmax)+1
				fmt.Fprintln(ans, a, b)
			}
		}
	}

	fmt.Fprintln(ans, 0, 0)

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}
