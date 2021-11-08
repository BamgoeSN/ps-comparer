package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var (
	nmin, nmax int   = 3, 6
	wmin, wmax int64 = 0, 10
)

func CreateInput() string {
	ans := new(strings.Builder)

	tc := RandRange(1, 1)
	for t := 0; t < tc; t++ {
		n := RandRange(nmin, nmax)
		uf := NewUnionFind(n)
		m := 0
		queries := make([][]string, 0)

		for uf.GroupNum != 1 {
			m++
			pick := RandRange(0, 1)
			switch pick {
			case 0:
				a, b := 0, 0
				for a == b {
					a = RandRange(0, n-1)
					b = RandRange(0, n-2)
					if b >= a {
						b++
					}
					if !uf.Differ(a, b) {
						b = a
					}
				}
				w := RandRange64(wmin, wmax)
				queries = append(queries, []string{"!", strconv.Itoa(a + 1), strconv.Itoa(b + 1), strconv.FormatInt(w, 10)})
				uf.Union(a, b)
			case 1:
				a, b := 0, 0
				for a == b {
					a = RandRange(0, n-1)
					b = RandRange(0, n-2)
					if b >= a {
						b++
					}
					if !uf.Differ(a, b) {
						b = a
					}
				}
				queries = append(queries, []string{"?", strconv.Itoa(a + 1), strconv.Itoa(b + 1)})
			}
		}

		m += n * (n - 1) / 2
		fmt.Fprintln(ans, n, m)
		for _, q := range queries {
			for _, v := range q {
				fmt.Fprintf(ans, "%v ", v)
			}
			fmt.Fprintln(ans)
		}
		for a := 0; a < n; a++ {
			for b := a + 1; b < n; b++ {
				fmt.Fprintf(ans, "? %d %d\n", a+1, b+1)
			}
		}
	}

	fmt.Fprintln(ans, 0, 0)

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}

func RandRange(l, r int) int       { return rand.Intn(r-l+1) + l }
func RandRange64(l, r int64) int64 { return rand.Int63n(r-l+1) + l }

//------------
// Union Find
//------------

type UnionFind struct {
	Parents  []int
	size     []int
	GroupNum int
}

// NewUnionFind creates new UnionFind instance
func NewUnionFind(size int) *UnionFind {
	uf := new(UnionFind)
	uf.Parents = make([]int, size)
	uf.size = make([]int, size)
	for i := range uf.Parents {
		uf.Parents[i] = -1
		uf.size[i] = 1
	}
	uf.GroupNum = size
	return uf
}

// find searches for the root of a subtree including a
func (uf *UnionFind) find(a int) int {
	if uf.Parents[a] == -1 {
		return a
	}
	list := make([]int, 0)
	for uf.Parents[a] != -1 {
		list = append(list, a)
		a = uf.Parents[a]
	}
	for _, v := range list {
		uf.Parents[v] = a
	}
	return a
}

// Differ determines if a and b are in different set
func (uf *UnionFind) Differ(a, b int) bool {
	return uf.find(a) != uf.find(b)
}

// Union merges two sets including a and b.
// Returns the same value with uf.Differ(a, b)
func (uf *UnionFind) Union(a, b int) bool {
	aRoot := uf.find(a)
	bRoot := uf.find(b)
	if aRoot != bRoot {
		uf.GroupNum--
		if uf.size[aRoot] < uf.size[bRoot] {
			uf.Parents[aRoot] = bRoot
		} else {
			uf.Parents[bRoot] = aRoot
		}
		uf.size[aRoot] = uf.size[aRoot] + uf.size[bRoot]
		uf.size[bRoot] = uf.size[aRoot]
	}
	return aRoot != bRoot
}

// GroupSize gets a size of a set which a is in
func (uf *UnionFind) GroupSize(a int) int {
	aRoot := uf.find(a)
	return uf.size[aRoot]
}
