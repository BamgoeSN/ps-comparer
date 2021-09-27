package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	nmin, nmax = 2, 6
)

type Edge struct {
	s, e int
	c    byte
}

func CreateInput() string {
	ans := new(strings.Builder)

	TC := 2

	for tc := 0; tc < TC; tc++ {
		n := rand.Intn(nmax-nmin+1) + nmin
		m := 0

		uf := NewUnionFind(n)
		graph := make([]Edge, 0)

		conrate := rand.Float64()
		redrate := rand.Float64()

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				conpick := rand.Float64()
				if conpick > conrate {
					continue
				}
				m++
				uf.Union(i, j)
				e := Edge{i, j, 'R'}
				exch := rand.Intn(2)
				if exch == 1 {
					e.s, e.e = e.e, e.s
				}

				redpick := rand.Float64()
				if redrate < redpick {
					e.c = 'B'
				}

				graph = append(graph, e)
			}
		}

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if !uf.Differ(i, j) {
					continue
				}
				m++
				e := Edge{i, j, 'R'}
				exch := rand.Intn(2)
				if exch == 1 {
					e.s, e.e = e.e, e.s
				}

				redpick := rand.Float64()
				if redrate < redpick {
					e.c = 'B'
				}

				graph = append(graph, e)
			}
		}

		k := rand.Intn(n)
		fmt.Fprintln(ans, n, m, k)
		for _, e := range graph {
			fmt.Fprintln(ans, string([]byte{e.c}), e.s+1, e.e+1)
		}
	}

	fmt.Fprintln(ans, 0, 0, 0)

	return TrimSpaces(ans.String())
}

func TrimSpaces(str string) string {
	flds := strings.Split(str, "\n")
	for i, v := range flds {
		flds[i] = strings.TrimSpace(v)
	}
	return strings.TrimSpace(strings.Join(flds, "\n"))
}

//------------
// Union Find
//------------

type UnionFind struct {
	Parents []int
	size    []int
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
