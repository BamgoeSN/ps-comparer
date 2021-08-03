package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmax, nmin int = 90000, 90000
)

func CreateInput() string {
	ans := new(strings.Builder)

	n := rand.Intn(nmax-nmin+1) + nmin
	fmt.Fprintln(ans, n)

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	// rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	for _, v := range arr {
		fmt.Fprintln(ans, v)
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
