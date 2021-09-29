package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax = 4, 4
	mmin, mmax = 1, 30
	smin, smax = 1, 5
	k, kmax    = 0, 3

	n, m  int
	fires []*Fire
)

type Fire struct{ r, c, m, s, d int }

func CreateInput() string {
	ans := new(strings.Builder)

	if k == 0 {
		n = RandRange(nmin, nmax)
		fires = make([]*Fire, 0)
		rate := rand.Float64() / 2
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				pick := rand.Float64()
				if pick < rate {
					m, d, s := RandRange(mmin, mmax), RandRange(0, 7), RandRange(smin, smax)
					fires = append(fires, &Fire{r + 1, c + 1, m, s, d})
				}
			}
		}
		m = len(fires)
		rand.Shuffle(len(fires), func(i, j int) { fires[i], fires[j] = fires[j], fires[i] })
	}
	k++

	fmt.Fprintln(ans, n, m, k)
	for _, f := range fires {
		fmt.Fprintln(ans, f.r, f.c, f.m, f.s, f.d)
	}

	if k == kmax {
		k = 0
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

func RandRange(l, r int) int { return rand.Intn(r-l+1) + l }
