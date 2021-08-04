package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var commands = []string{"push", "pop", "size", "empty", "top"}

var (
	nmin, nmax int   = 1, 5
	vmin, vmax int64 = 1, 10
)

func CreateInput() string {
	ans := new(strings.Builder)

	N := rand.Intn(nmax-nmin+1) + nmin
	fmt.Fprintln(ans, N)

	for i := 0; i < N; i++ {
		q := commands[rand.Intn(len(commands))]
		switch q {
		case "push":
			fmt.Fprintln(ans, q, rand.Int63n(vmax-vmin+1)+vmin)
		default:
			fmt.Fprintln(ans, q)
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
