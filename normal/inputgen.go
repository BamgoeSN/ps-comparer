package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CreateInput() string {
	ans := new(strings.Builder)

	nmax := 26
	lenmax := 100000

	N := rand.Intn(nmax-1) + 2
	fmt.Fprintln(ans, N)

	length := rand.Intn(lenmax) + 1
	for i := 0; i < length; i++ {
		ans.WriteByte('a' + byte(rand.Intn(nmax)))
	}

	return TrimSpaces(ans.String())
}
