package algoritmos

import (
	"fmt"
)

func NaivLoopUnrollingThree(A, B, Result [][]float64, N, P, M int) {
	if P%3 == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux float64
				for k := 0; k < P; k += 3 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j]
				}
				Result[i][j] = aux
			}
		}
	} else if P%3 == 1 {
		var PP int = P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux float64
				for k := 0; k < PP; k += 3 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j]
			}
		}
	} else {
		var PP int = P - 2
		var PPP int = P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux float64
				for k := 0; k < PP; k += 3 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j]
			}
		}
	}
}

func LlamarNaivLoopUnrollingThree() {
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

	// Compute the matrix product using the NaivLoopUnrollingThree algorithm
	NaivLoopUnrollingThree(A, B, Result, 3, 3, 3)

	// Print the result matrix
	fmt.Println("Result:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%f ", Result[i][j])
		}
		fmt.Println()
	}
}
