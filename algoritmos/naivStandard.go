package algoritmos

import "fmt"

func NaivStandard(A [][]float64, B [][]float64, Result [][]float64, N, P, M int) {
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            aux := 0.0
            for k := 0; k < P; k++ {
                aux += A[i][k] * B[k][j]
            }
            Result[i][j] = aux
        }
    }
}

func LlamarNaivStandart() {
    // Definir las dimensiones de las matrices
    N := 2
    P := 3
    M := 2

    // Declarar y asignar valores a las matrices A y B
    A := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
    B := [][]float64{{7.0, 8.0}, {9.0, 10.0}, {11.0, 12.0}}

    // Crear una matriz para el resultado
    Result := make([][]float64, N)
    for i := range Result {
        Result[i] = make([]float64, M)
    }

    // Llamar a la función para multiplicar las matrices
    NaivStandard(A, B, Result, N, P, M)

    // Imprimir la matriz resultante
    fmt.Println("Matriz Resultante:")
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            fmt.Printf("%.2f ", Result[i][j])
        }
        fmt.Println()
    }
}