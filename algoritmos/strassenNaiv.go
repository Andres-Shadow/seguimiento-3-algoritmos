package algoritmos

import (
	"fmt"
	"math"
)

func StrassenNaiv(A [][]float64, B [][]float64, Result [][]float64, N int, P int, M int) {
	MaxSize := math.Max(float64(N), float64(P))
	MaxSize = math.Max(MaxSize, float64(M))
	if MaxSize < 16 {
		MaxSize = 16 // otherwise it is not possible to compute k
	}
	k := int(math.Floor(math.Log2(MaxSize)/math.Log2(2))) - 4
	m := int(math.Floor(MaxSize*math.Pow(2, -float64(k)))) + 1
	NewSize := m * int(math.Pow(2, float64(k)))

	// add zero rows and columns to use Strassens algorithm
	NewA := make([][]float64, NewSize)
	NewB := make([][]float64, NewSize)
	AuxResult := make([][]float64, NewSize)
	for i := 0; i < NewSize; i++ {
		NewA[i] = make([]float64, NewSize)
		NewB[i] = make([]float64, NewSize)
		AuxResult[i] = make([]float64, NewSize)
	}

	for i := 0; i < NewSize; i++ {
		for j := 0; j < NewSize; j++ {
			NewA[i][j] = 0.0
			NewB[i][j] = 0.0
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			NewA[i][j] = A[i][j]
		}
	}

	for i := 0; i < P; i++ {
		for j := 0; j < M; j++ {
			NewB[i][j] = B[i][j]
		}
	}

	StrassenNaivStep(NewA, NewB, AuxResult, NewSize, m)

	// extract the result
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = AuxResult[i][j]
		}
	}
}

func StrassenNaivStep(A [][]float64, B [][]float64, Result [][]float64, Size int, m int) {
	// implementation of Strassen's algorithm
	// ...
}

func LlamarStrassenNaiv(opcion int) {
	// Example usage
	N := 2
	P := 2
	M := 2
	A := [][]float64{{1, 2}, {3, 4}}
	B := [][]float64{{5, 6}, {7, 8}}
	Result := make([][]float64, N)
	for i := 0; i < N; i++ {
		Result[i] = make([]float64, M)
	}

	StrassenNaiv(A, B, Result, N, P, M)

	fmt.Println(Result)
}
