package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
	"time"
)

func NaivLoopUnrollingTwo(A, B, Result [][]int, N, P, M int) {
	if P%2 == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
				for k := 0; k < P; k += 2 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
				}
				Result[i][j] = aux
			}
		}
	} else {
		var PP int = P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
				for k := 0; k < PP; k += 2 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j]
			}
		}
	}
}

func LlamarNaivLoopUnrollingTwo(opcion int) {
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
	// Compute the matrix product using the NaivLoopUnrollingTwo algorithm
	NaivLoopUnrollingTwo(A, B, Result, N, N, N)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
