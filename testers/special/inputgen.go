package main

import (
	"strings"
)

var (
	k = 1
)

func CreateInput() string {
	ans := new(strings.Builder)

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}
