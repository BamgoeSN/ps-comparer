package main

import (
	"math/rand"
	"strings"
)

var list = []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z=", "c", "d", "z", "l", "j", "n", "s", "a", "g", "dz"}

func CreateInput() string {
	ans := new(strings.Builder)

	lenmax := 20
	length := rand.Intn(lenmax) + 1
	for ans.Len() < length {
		b := list[rand.Intn(len(list))]
		ans.WriteString(b)
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
