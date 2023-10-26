package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
	"time"
)

func NaivKahan(A, B, Result [][]int, N, P, M int) {
	var t, sum, err int
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

func LlamarNaivKahan(opcion int) {

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
	// Compute the matrix product using the NaivKahan algorithm
	NaivKahan(A, B, Result, N, N, N)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)

}
