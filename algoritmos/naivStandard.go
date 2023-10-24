package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
	"time"
)

func NaivStandard(A [][]int, B [][]int, Result [][]int, N, P, M int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			aux := 0
			for k := 0; k < P; k++ {
				aux += A[i][k] * B[k][j]
			}
			Result[i][j] = aux
		}
	}
}

func LlamarNaivStandart(opcion int) {

	startTime := time.Now()

	// Definir las dimensiones de las matrices
	var nombre1, nombre2 string
	var N, P, M int

	switch opcion {
	case 1:
		N = 1024
		P = 1024
		M = 1024
		nombre1 = "datos.dat"
		nombre2 = "datos2.dat"
		break
	case 2:
		N = 2048
		P = 2048
		M = 2048
		nombre1 = "datos3.dat"
		nombre2 = "datos4.dat"
		break
	case 3:
		N = 4096
		P = 4096
		M = 4096
		nombre1 = "datos1.dat"
		nombre2 = "datos2.dat"
		break
	}

	// Declarar y asignar valores a las matrices A y B
	A, _ := herramientas.LoadMatrixFromFile(nombre1, N, N)
	B, _ := herramientas.LoadMatrixFromFile(nombre2, N, N)

	// Crear una matriz para el resultado
	Result := make([][]int, N)
	for i := range Result {
		Result[i] = make([]int, M)
	}

	// Llamar a la función para multiplicar las matrices
	NaivStandard(A, B, Result, N, P, M)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
