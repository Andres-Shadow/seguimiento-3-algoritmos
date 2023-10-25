package algoritmos

import (
	"fmt"
	"math"
	"seg_3_algoritmos/herramientas"
	"time"
)

func StrassenNaiv(A [][]int, B [][]int, Result [][]int, N int, P int, M int) {
	MaxSize := math.Max(float64(int(N)), float64(int(P)))
	MaxSize = math.Max(MaxSize, float64(int(M)))
	if MaxSize < 16 {
		MaxSize = 16 // otherwise it is not possible to compute k
	}
	k := int(math.Floor(math.Log2(MaxSize)/math.Log2(2))) - 4
	m := int(math.Floor(MaxSize*math.Pow(2, float64(-int(k))))) + 1
	NewSize := m * int(math.Pow(2, float64(int(k))))

	// add zero rows and columns to use Strassens algorithm
	NewA := make([][]int, NewSize)
	NewB := make([][]int, NewSize)
	AuxResult := make([][]int, NewSize)
	for i := 0; i < NewSize; i++ {
		NewA[i] = make([]int, NewSize)
		NewB[i] = make([]int, NewSize)
		AuxResult[i] = make([]int, NewSize)
	}

	for i := 0; i < NewSize; i++ {
		for j := 0; j < NewSize; j++ {
			NewA[i][j] = 0.0
			NewB[i][j] = 0.0
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			NewA[i][j] = A[i][j]
		}
	}

	for i := 0; i < P; i++ {
		for j := 0; j < M; j++ {
			NewB[i][j] = B[i][j]
		}
	}

	StrassenNaivStep(NewA, NewB, AuxResult, NewSize, m)

	// extract the result
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = AuxResult[i][j]
		}
	}
}

func StrassenNaivStep(A, B, Result [][]int, N, m int) {
	if N%2 == 0 && N > m {
		NewSize := N / 2

		A11 := make([][]int, NewSize)
		A12 := make([][]int, NewSize)
		A21 := make([][]int, NewSize)
		A22 := make([][]int, NewSize)
		B11 := make([][]int, NewSize)
		B12 := make([][]int, NewSize)
		B21 := make([][]int, NewSize)
		B22 := make([][]int, NewSize)
		ResultPart11 := make([][]int, NewSize)
		ResultPart12 := make([][]int, NewSize)
		ResultPart21 := make([][]int, NewSize)
		ResultPart22 := make([][]int, NewSize)
		Helper1 := make([][]int, NewSize)
		Helper2 := make([][]int, NewSize)
		Aux1 := make([][]int, NewSize)
		Aux2 := make([][]int, NewSize)
		Aux3 := make([][]int, NewSize)
		Aux4 := make([][]int, NewSize)
		Aux5 := make([][]int, NewSize)
		Aux6 := make([][]int, NewSize)
		Aux7 := make([][]int, NewSize)

		for i := 0; i < NewSize; i++ {
			A11[i] = make([]int, NewSize)
			A12[i] = make([]int, NewSize)
			A21[i] = make([]int, NewSize)
			A22[i] = make([]int, NewSize)
			B11[i] = make([]int, NewSize)
			B12[i] = make([]int, NewSize)
			B21[i] = make([]int, NewSize)
			B22[i] = make([]int, NewSize)
			ResultPart11[i] = make([]int, NewSize)
			ResultPart12[i] = make([]int, NewSize)
			ResultPart21[i] = make([]int, NewSize)
			ResultPart22[i] = make([]int, NewSize)
			Helper1[i] = make([]int, NewSize)
			Helper2[i] = make([]int, NewSize)
			Aux1[i] = make([]int, NewSize)
			Aux2[i] = make([]int, NewSize)
			Aux3[i] = make([]int, NewSize)
			Aux4[i] = make([]int, NewSize)
			Aux5[i] = make([]int, NewSize)
			Aux6[i] = make([]int, NewSize)
			Aux7[i] = make([]int, NewSize)
		}

		for i := 0; i < NewSize; i++ {
			for j := 0; j < NewSize; j++ {
				A11[i][j] = A[i][j]
				A12[i][j] = A[i][NewSize+j]
				A21[i][j] = A[NewSize+i][j]
				A22[i][j] = A[NewSize+i][NewSize+j]
				B11[i][j] = B[i][j]
				B12[i][j] = B[i][NewSize+j]
				B21[i][j] = B[NewSize+i][j]
				B22[i][j] = B[NewSize+i][NewSize+j]
			}
		}

		// Computing the seven auxiliary variables
		Plus(A11, A22, Helper1, NewSize, NewSize)
		Plus(B11, B22, Helper2, NewSize, NewSize)
		StrassenNaivStep(Helper1, Helper2, Aux1, NewSize, m)
		Plus(A21, A22, Helper1, NewSize, NewSize)
		StrassenNaivStep(Helper1, B11, Aux2, NewSize, m)
		Minus(B12, B22, Helper1, NewSize, NewSize)
		StrassenNaivStep(A11, Helper1, Aux3, NewSize, m)
		Minus(B21, B11, Helper1, NewSize, NewSize)
		StrassenNaivStep(A22, Helper1, Aux4, NewSize, m)
		Plus(A11, A12, Helper1, NewSize, NewSize)
		StrassenNaivStep(Helper1, B22, Aux5, NewSize, m)
		Minus(A21, A11, Helper1, NewSize, NewSize)
		Plus(B11, B12, Helper2, NewSize, NewSize)
		StrassenNaivStep(Helper1, Helper2, Aux6, NewSize, m)
		Minus(A12, A22, Helper1, NewSize, NewSize)
		Plus(B21, B22, Helper2, NewSize, NewSize)
		StrassenNaivStep(Helper1, Helper2, Aux7, NewSize, m)

		// Computing the four parts of the result
		Plus(Aux1, Aux4, ResultPart11, NewSize, NewSize)
		Minus(ResultPart11, Aux5, ResultPart11, NewSize, NewSize)
		Plus(ResultPart11, Aux7, ResultPart11, NewSize, NewSize)
		Plus(Aux3, Aux5, ResultPart12, NewSize, NewSize)
		Plus(Aux2, Aux4, ResultPart21, NewSize, NewSize)
		Plus(Aux1, Aux3, ResultPart22, NewSize, NewSize)
		Minus(ResultPart22, Aux2, ResultPart22, NewSize, NewSize)
		Plus(ResultPart22, Aux6, ResultPart22, NewSize, NewSize)

		// Store results in the "result matrix"
		for i := 0; i < NewSize; i++ {
			for j := 0; j < NewSize; j++ {
				Result[i][j] = ResultPart11[i][j]
				Result[i][NewSize+j] = ResultPart12[i][j]
				Result[NewSize+i][j] = ResultPart21[i][j]
				Result[NewSize+i][NewSize+j] = ResultPart22[i][j]
			}
		}

		// Free helper variables

	} else {
		// Use the naive algorithm
		NaivStandard(A, B, Result, N, N, N)
	}
}

func LlamarStrassenNaiv(opcion int) {
	// Example usage
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
		nombre1 = "datos5.dat"
		nombre2 = "datos6.dat"
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

	StrassenNaiv(A, B, Result, N, P, M)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}

func Plus(A, B, Result [][]int, N, M int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = A[i][j] + B[i][j]
		}
	}
}

func Minus(A, B, Result [][]int, N, M int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = A[i][j] - B[i][j]
		}
	}
}
