package main

import (
	"math/big"
	"strconv"
	"strings"
)

type Input struct {
	base int
	str  string
}

func Judge(in, out string) bool {
	// Parse input
	inFields := strings.Fields(in)[1:]
	inputs := make([]Input, len(inFields)/2)
	for i := range inputs {
		inputs[i].base, _ = strconv.Atoi(inFields[2*i])
		inputs[i].str = swapCase(inFields[2*i+1])
	}

	// Parse output
	outFields := strings.Fields(out)
	outputs := make([][3]string, len(outFields)/3)
	if len(outputs)*3 != len(outFields) {
		return false
	}
	for i := range outputs {
		for j := 0; j < 3; j++ {
			outputs[i][j] = swapCase(outFields[3*i+j])
		}
	}

	if len(inputs) != len(outputs) {
		return false
	}

	// Check if every sums are proper
	for i := range inputs {
		base := inputs[i].base
		target, _ := new(big.Int).SetString(inputs[i].str, base)
		collect := make([]*big.Int, 3)
		for j := range collect {
			tmp, check := new(big.Int).SetString(outputs[i][j], base)
			if !check {
				return false
			}
			collect[j] = tmp
		}
		please := big.NewInt(0)
		for _, v := range collect {
			please.Add(please, v)
		}

		if target.Cmp(please) != 0 {
			return false
		}
	}
	return true
}

func swapCase(str string) string {
	bts := []byte(str)
	res := make([]byte, len(bts))
	for i, v := range bts {
		switch {
		case v >= '0' && v <= '9':
			res[i] = v
		case v >= 'A' && v <= 'Z':
			res[i] = v + 'a' - 'A'
		case v >= 'a' && v <= 'z':
			res[i] = v + 'A' - 'a'
		}
	}
	return string(res)
}
