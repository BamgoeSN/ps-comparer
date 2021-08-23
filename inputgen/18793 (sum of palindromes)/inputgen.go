package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax int = 5, 62
	lmin, lmax int = 1, 1000000
	lookup         = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func CreateInput() string {
	ans := new(strings.Builder)

	TC := 1
	fmt.Fprintln(ans, TC)
	for tc := 0; tc < TC; tc++ {
		base := rand.Intn(nmax-nmin+1) + nmin
		length := rand.Intn(lmax-lmin+1) + lmin

		output := make([]byte, length)
		output[0] = lookup[rand.Intn(base-1)+1]

		for i := 1; i < length; i++ {
			output[i] = lookup[rand.Intn(base)]
		}

		fmt.Fprintln(ans, base, string(output))
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
