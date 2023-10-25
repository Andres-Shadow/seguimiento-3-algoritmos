package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
	"time"
)

func NaivOnArray(A, B, Result [][]int, N, P, M int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = 0.0
			for k := 0; k < P; k++ {
				Result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func LlamarNaivOnArray(opcion int) {

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

	var Result [][]int = make([][]int, 3)

	// Compute the matrix product using the NaivOnArray algorithm
	NaivOnArray(A, B, Result, 3, 3, 3)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
