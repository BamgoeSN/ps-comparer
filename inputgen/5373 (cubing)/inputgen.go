package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var sides = []string{"U", "D", "F", "B", "L", "R"}
var dir = []string{"+", "-"}

func CreateInput() string {
	ans := new(strings.Builder)
	T := 1
	fmt.Fprintln(ans, T)

	for t := 0; t < T; t++ {
		try := rand.Intn(10) + 1
		fmt.Fprintln(ans, try)
		instr := make([]string, try)
		for i := range instr {
			instr[i] = sides[rand.Intn(6)] + dir[rand.Intn(2)]
		}
		for _, v := range instr[:len(instr)-1] {
			fmt.Fprintf(ans, "%s ", v)
		}
		fmt.Fprintln(ans, instr[len(instr)-1])
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
