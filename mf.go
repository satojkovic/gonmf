package main

import (
	"math"

	"github.com/skelterjohn/go.matrix"
)

type Params struct {
	Steps int
	Alpha float64
	Beta  float64
}

func matrixFactorization(R [][]int, P [][]float64, Q [][]float64, K int, p Params) ([][]float64, [][]float64) {
	nU, nD := len(R), len(R[0])

	Pm := matrix.MakeDenseMatrixStacked(P)
	Qm := matrix.MakeDenseMatrixStacked(Q)
	Qm = Qm.Transpose()

	for step := 0; step < p.Steps; step++ {
		for i := 0; i < nU; i++ {
			for j := 0; j < nD; j++ {
				if R[i][j] > 0.0 {
					Pmi := Pm.GetRowVector(i)
					Qmj := Qm.GetColVector(j)
					dotp, _ := Pmi.TimesDense(Qmj)
					eij := float64(R[i][j]) - dotp.Get(0, 0)

					for k := 0; k < K; k++ {
						pik := Pm.Get(i, k)
						qkj := Qm.Get(k, j)
						Pm.Set(i, k, pik+p.Alpha*(2*eij*qkj-p.Beta*pik))
						Qm.Set(k, j, qkj+p.Alpha*(2*eij*pik-p.Beta*qkj))
					}
				}
			}
		}
		e := 0.0
		for i := 0; i < nU; i++ {
			for j := 0; j < nD; j++ {
				if R[i][j] > 0.0 {
					Pmi := Pm.GetRowVector(i)
					Qmj := Qm.GetColVector(j)
					dotp, _ := Pmi.TimesDense(Qmj)
					e = e + math.Pow(dotp.Get(0, 0), 2.0)

					for k := 0; k < K; k++ {
						e = e + (p.Beta/2.0)*(math.Pow(Pm.Get(i, k), 2)+math.Pow(Qm.Get(k, j), 2))
					}
				}
			}
		}
		if e < 0.001 {
			break
		}
	}

	return Pm.Arrays(), Qm.Transpose().Arrays()
}
