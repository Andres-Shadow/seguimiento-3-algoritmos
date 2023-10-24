package algoritmos

import (
	"fmt"
	"seg_3_algoritmos/herramientas"
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

func LlamarNaivStandart() {
	// Definir las dimensiones de las matrices
	N := 1024
	P := 1024
	M := 1024

	// Declarar y asignar valores a las matrices A y B
	A, _ := herramientas.LoadMatrixFromFile("datos.dat", 1024, 1024)
	B, _ := herramientas.LoadMatrixFromFile("datos2.dat", 1024, 1024)

	// Crear una matriz para el resultado
	Result := make([][]int, N)
	for i := range Result {
		Result[i] = make([]int, M)
	}

	// Llamar a la función para multiplicar las matrices
	NaivStandard(A, B, Result, N, P, M)

	// Imprimir la matriz resultante
	fmt.Println("Matriz Resultante:")
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Print(Result[i][j], " ")
		}
		fmt.Println()
	}
}
