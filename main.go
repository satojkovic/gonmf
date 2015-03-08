package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	R := [][]int{
		{5, 3, 0, 1},
		{4, 0, 0, 1},
		{1, 1, 0, 5},
		{1, 0, 0, 4},
		{0, 1, 5, 4},
	}

	N := len(R)
	M := len(R[0])
	K := 2

	rand.Seed(time.Now().Unix())
	P := randMat(N, K)
	Q := randMat(M, K)

	nP, nQ := matrixFactorization(
		R, P, Q, K,
		Params{Steps: 5000, Alpha: 0.002, Beta: 0.02},
	)

	fmt.Printf("%v\n", nP)
	fmt.Printf("%v\n", nQ)
}

func randMat(row, col int) [][]float64 {
	rmat := make([][]float64, row)
	for r := 0; r < row; r++ {
		cols := make([]float64, col)
		for c := 0; c < col; c++ {
			cols[c] = rand.Float64()
		}
		rmat[r] = cols
	}

	return rmat
}
