package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var ()

func CreateInput() string {
	ans := new(strings.Builder)

	n, c := 1000, 1000
	strlist := make([]string, n+c)
	fmt.Fprintln(ans, n, c)
	for i := range strlist {
		str := make([]byte, 1000)
		for i := range str {
			str[i] = 'a' + byte(rand.Intn(26))
		}
		strlist[i] = string(str)
		fmt.Fprintln(ans, strlist[i])
	}

	rate := rand.Float64()
	q := 20000
	fmt.Fprintln(ans, q)
	for i := 0; i < q; i++ {
		pick := rand.Float64()
		if pick < rate {
			var str = []string{}
			str = append(str, strlist[rand.Intn(c)])
			str = append(str, strlist[rand.Intn(n)+c])
			fmt.Fprintln(ans, str[0]+str[1])
		} else {
			str := make([]byte, 2000)
			for i := range str {
				str[i] = 'a' + byte(rand.Intn(26))
			}
			fmt.Fprintln(ans, string(str))
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
