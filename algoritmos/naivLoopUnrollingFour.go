package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
	"time"
)

func NaivLoopUnrollingFour(A, B, Result [][]int, N, P, M int) {
	if P%4 == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
				for k := 0; k < P; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux
			}
		}
	} else if P%4 == 1 {
		var PP int = P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
				for k := 0; k < PP; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j]
			}
		}
	} else if P%4 == 2 {
		var PP int = P - 2
		var PPP int = P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
				for k := 0; k < PP; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j]
			}
		}
	} else {
		var PP int = P - 3
		var PPP int = P - 2
		var PPPP int = P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				var aux int
				for k := 0; k < PP; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j] + A[i][PPPP]*B[PPPP][j]
			}
		}
	}
}

func LlamarNaivLoopUnrollingFour(opcion int) {
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

	NaivLoopUnrollingFour(A, B, Result, N, N, N)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
