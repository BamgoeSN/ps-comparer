package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var (
	nmin, nmax int = 2, 10
)

func CreateInput() string {
	ans := new(strings.Builder)

	N := rand.Intn(nmax-nmin+1) + nmin
	fmt.Fprintln(ans, N)
	arr := make([][]int, N)
	for r := range arr {
		arr[r] = make([]int, N)
	}

	rate := make([]float64, 6)
	for i := range rate {
		rate[i] = rand.Float64()
	}
	sort.Float64s(rate)

	for r := range arr {
		for c := range arr[r] {
			pick := rand.Float64()
			arr[r][c] = sort.Search(len(rate), func(i int) bool { return pick < rate[i] })
		}
	}

	sr, sc := rand.Intn(N), rand.Intn(N)
	arr[sr][sc] = 9

	for _, r := range arr {
		for _, e := range r {
			fmt.Fprintf(ans, "%d ", e)
		}
		fmt.Fprintln(ans)
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
