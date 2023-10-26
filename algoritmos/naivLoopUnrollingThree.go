package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
	"time"
)

func NaivLoopUnrollingThree(A, B, Result [][]int, N, P, M int) {
	if P%3 == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
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
				var aux int
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
				var aux int
				for k := 0; k < PP; k += 3 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j]
			}
		}
	}
}

func LlamarNaivLoopUnrollingThree(opcion int) {
	startTime := time.Now()

	// Definir las dimensiones de las matrices
	var nombre1, nombre2 string
	var N int

	switch opcion {
	case 1:
		N = 1024
		nombre1 = "datos.dat"
		nombre2 = "datos2.dat"
		break
	case 2:
		N = 2048
		nombre1 = "datos3.dat"
		nombre2 = "datos4.dat"
		break
	case 3:
		N = 4096
		nombre1 = "datos5.dat"
		nombre2 = "datos6.dat"
		break
	}

	// Declarar y asignar valores a las matrices A y B
	A, _ := herramientas.LoadMatrixFromFile(nombre1, N, N)
	B, _ := herramientas.LoadMatrixFromFile(nombre2, N, N)

	var Result [][]int = make([][]int, N)
	for i := range Result {
		Result[i] = make([]int, N)
	}
	// Compute the matrix product using the NaivLoopUnrollingThree algorithm
	NaivLoopUnrollingThree(A, B, Result, N, N, N)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}