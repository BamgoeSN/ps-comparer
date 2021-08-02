package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//------
// Main
//------

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	N := nextInt()
	roads := make([]int64, N-1)
	cities := make([]int64, N)
	for i := range roads {
		roads[i] = nextInt64()
	}
	for i := range cities {
		cities[i] = nextInt64()
	}

	var min int64 = math.MaxInt64
	var cost int64 = 0
	for i, v := range roads {
		if cities[i] < min {
			min = cities[i]
		}
		cost += min * v
	}
	fmt.Fprintln(wr, cost)
}

//---------
// Fast IO
//---------

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() (res int) {
	sc.Scan()
	text := sc.Text()
	v, _ := strconv.Atoi(text)
	return v
}

func nextInt64() (res int64) {
	sc.Scan()
	text := sc.Text()
	v, _ := strconv.ParseInt(text, 10, 64)
	return v
}
