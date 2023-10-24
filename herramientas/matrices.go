package herramientas

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
)

func GenearMatricesPorTamanio(tam int, nombre, nombre2 string) {
	switch tam {
	case 1:
		tam = 1024
		break
	case 2:
		tam = 2048
		break
	case 3:
		tam = 4096
		break
	}

	m := GenerateRandomMatrix(tam, tam)
	SaveMatrixToFile(m, nombre)
	fmt.Println("se ha generado la matriz 1 de tamaño ", tam)
	m = GenerateRandomMatrix(tam, tam)
	SaveMatrixToFile(m, nombre2)
	fmt.Println("se ha generado la matriz 1 de tamaño ", tam)
}

func GenerateRandomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100) // Genera números aleatorios entre 0 y 99
		}
	}
	return matrix
}

func SaveMatrixToFile(matrix [][]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range matrix {
		for _, value := range row {
			err := binary.Write(file, binary.LittleEndian, int32(value))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func LoadMatrixFromFile(filename string, rows, cols int) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			var value int32
			err := binary.Read(file, binary.LittleEndian, &value)
			if err != nil {
				return nil, err
			}
			matrix[i][j] = int(value)
		}
	}

	return matrix, nil
}
