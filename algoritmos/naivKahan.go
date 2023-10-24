package algoritmos

import (
	"fmt"
)

func NaivKahan(A, B, Result [][]float64, N, P, M int) {
	var t, sum, err float64
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			sum = 0.0
			err = 0.0
			for k := 0; k < P; k++ {
				err += A[i][k] * B[k][j]
				t = sum + err
				err = (sum - t) + err
				sum = t
			}
			Result[i][j] = sum
		}
	}
}

func LlamarNaivKahan() {
	var A, B, Result [][]float64 = make([][]float64, 3), make([][]float64, 3), make([][]float64, 3)
	for i := 0; i < 3; i++ {
		A[i] = make([]float64, 3)
		B[i] = make([]float64, 3)
		Result[i] = make([]float64, 3)
	}

	// Initialize the matrices A and B
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			A[i][j] = float64(i + j)
			B[i][j] = float64(2*i + j)
		}
	}

	// Compute the matrix product using the NaivKahan algorithm
	NaivKahan(A, B, Result, 3, 3, 3)

	// Print the result matrix
	fmt.Println("Result:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%f ", Result[i][j])
		}
		fmt.Println()
	}
}
